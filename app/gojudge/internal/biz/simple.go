package biz

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sastoj/ent"
	"sastoj/ent/problem"
	"sastoj/pkg/mq"
	u "sastoj/pkg/util"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// Simple Adapter
type Simple struct {
	middleware *Middleware
	config     *u.JudgeConfig
}

// handleSubmit process the submission
// 1. get problem, problem-cases and config-file and by problem-id
// 2. compile submission
// 2. get in-out by problem-id and case-id
// 3. judge each case-submission and save, then create submission and save
func (s *Simple) handleSubmit(v *Submit) error {
	ctx := context.Background()

	redisKey := fmt.Sprintf("submission:%d:%s", v.UserID, v.ID)

	submission := &mq.Submission{
		ID:         v.ID,
		UserID:     v.UserID,
		ProblemID:  v.ProblemID,
		Code:       v.Code,
		Status:     v.Status,
		Point:      0,
		CreateTime: v.CreateTime,
		TotalTime:  0,
		MaxMemory:  0,
		Language:   v.Language,
		CompileMsg: "",
		Token:      v.Token,
	}

	//get problem from ent
	pro, err := s.middleware.ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		WithProblemCases().
		First(ctx)
	if err != nil {
		submission.Status = u.SystemError
		submission.CompileMsg = "problem not found"
		marshal, _ := json.Marshal(submission)
		s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour)
		return errors.New("problem not found")
	}

	// TODO: cache and refresh config by problem-id

	//compile submit
	fileID, result, err := s.middleware.goJudge.Compile([]byte(v.Code), v.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			submission.Status = u.CompileError
			stderr := "Compile error"
			compileError, ok := result.Files["stderr"]
			if ok {
				stderr = string(compileError)
			}
			save, err := s.middleware.ent.Submission.Create().
				SetCaseVersion(int8(pro.CaseVersion)).
				SetUsersID(v.UserID).
				SetCode(v.Code).
				SetLanguage(v.Language).
				SetPoint(0).
				SetStatus(u.CompileError).
				SetCreateTime(v.CreateTime).
				SetTotalTime(0).
				SetMaxMemory(0).
				SetProblemsID(v.ProblemID).
				SetCompileMessage(stderr).
				Save(ctx)
			if err != nil {
				log.Errorf("save error: %v", err)
			}
			submission.ID = strconv.FormatInt(save.ID, 10)
			submission.CompileMsg = stderr
		} else {
			submission.Status = u.SystemError
			submission.CompileMsg = err.Error()
		}
		marshal, _ := json.Marshal(submission)
		s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour)
		return err
	}

	var totalTime uint64 = 0
	var maxMemory uint64 = 0
	var score int16 = 0
	var status int16 = u.Accepted
	var builders = make([]*ent.SubmissionCaseCreate, 0)

	//test each case
	//TODO index
	for _, c := range s.config.Task.Cases {

		fileIndex, err := strconv.Atoi(strings.Split(c.Input, ".")[0])
		if err != nil {
			return err
		}

		in, ans, err := s.middleware.fileManage.FetchCase(v.ProblemID, c.Input, c.Answer)
		//TODO remove
		in = u.RemoveCr(in)
		ans = u.RemoveCr(ans)
		if err != nil {
			_ = s.middleware.goJudge.DeleteFile(fileID)
			return err
		}
		result, err := s.middleware.goJudge.Judge(in, v.Language, fileID, uuid.NewString(), uint64(s.config.ResourceLimits.Time), uint64(s.config.ResourceLimits.Time*2), uint64(s.config.ResourceLimits.Memory), int64(len(ans)))
		if err != nil {
			_ = s.middleware.goJudge.DeleteFile(fileID)
			s.middleware.logger.Errorf("judge error: %v", err)
		}
		s.middleware.logger.Infof("submission: %s result: %+v", v.ID, result)
		state := u.FromGoJudgeState(int32(result.Status.Number()))

		if out := result.Files["stdout"]; state == u.Accepted && out != nil {
			if pro.RestrictPresentation && !bytes.Equal(out, ans) {
				if u.BytesMatchIgnoringSpacesAndNewlines(out, ans) {
					state = u.PresentationError
				} else {
					state = u.WrongAnswer
				}
			} else if !u.BytesMatchIgnoringSpacesAndNewlines(out, ans) {
				state = u.WrongAnswer
			}
		}

		if state != u.Accepted {
			if status == u.Accepted {
				status = state
			}
			//get problem cases ID
			var problemCasesID int64
			for _, problemCase := range pro.Edges.ProblemCases {
				if problemCase.Index == int16(fileIndex) {
					problemCasesID = problemCase.ID
				}
			}
			create := s.middleware.ent.SubmissionCase.Create().
				SetProblemCasesID(problemCasesID).
				SetMemory(int32(result.Memory)).
				SetTime(int32(result.RunTime)).
				SetState(state).
				SetPoint(0)
			msg, ok := result.Files["stderr"]
			if ok {
				create.SetMessage(string(msg))
			} else {
				create.SetMessage(result.Error)
			}
			builders = append(builders, create)
			//TODO if oi break
			continue
		}

		//Accept
		totalTime += result.RunTime
		if maxMemory < result.Memory {
			maxMemory = result.Memory
		}
		score += c.Score
		//get problem cases ID
		var problemCasesID int64
		for _, problemCase := range pro.Edges.ProblemCases {
			if problemCase.Index == int16(fileIndex) {
				problemCasesID = problemCase.ID
			}
		}

		//gen builder
		create := s.middleware.ent.SubmissionCase.Create().
			SetProblemCasesID(problemCasesID).
			SetMemory(int32(result.Memory)).
			SetTime(int32(result.RunTime)).
			SetMessage(result.Status.String()).
			SetState(state).
			SetPoint(score)
		builders = append(builders, create)
	}

	err = s.middleware.goJudge.DeleteFile(fileID)
	if err != nil {
		s.middleware.logger.Errorf("delete gojudge file error: %v", err)
	}

	//Save into database
	save, err := s.middleware.ent.Submission.Create().
		SetCaseVersion(int8(pro.CaseVersion)).
		SetUsersID(v.UserID).
		SetCode(v.Code).
		SetLanguage(v.Language).
		SetPoint(score).
		SetStatus(status).
		SetCreateTime(v.CreateTime).
		SetTotalTime(int32(totalTime)).
		SetMaxMemory(int32(maxMemory)).
		SetProblemsID(v.ProblemID).
		Save(context.Background())
	if err != nil {
		s.middleware.logger.Errorf("save error: %v", err)
	}

	submission.ID = strconv.FormatInt(save.ID, 10)
	submission.Status = status
	submission.Point = score
	submission.TotalTime = int32(totalTime)
	submission.MaxMemory = int32(maxMemory)
	marshal, _ := json.Marshal(submission)
	s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour)

	for _, build := range builders {
		build.SetSubmissionID(save.ID)
	}
	_, err = s.middleware.ent.SubmissionCase.CreateBulk(builders...).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
