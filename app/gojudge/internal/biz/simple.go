package biz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sastoj/ent"
	"sastoj/ent/problem"
	"sastoj/pkg/mq"
	u "sastoj/pkg/util"
	"strconv"
	"strings"
	"time"

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

	//get problem from ent
	pro, err := s.middleware.ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		WithProblemCases().
		First(ctx)
	if err != nil {
		return err
	}

	submission, err := s.middleware.ent.Submission.Create().
		SetCode(v.Code).
		SetStatus(u.Waiting).
		SetCompileMessage("").
		SetPoint(0).
		SetCreateTime(v.CreateTime).
		SetTotalTime(0).
		SetMaxMemory(0).
		SetLanguage(v.Language).
		SetCaseVersion(int8(pro.CaseVersion)).
		SetProblemsID(v.ProblemID).
		SetUsersID(v.UserID).
		Save(ctx)
	if err != nil {
		return err
	}
	redisKey := fmt.Sprintf("submission:%d:%s", v.UserID, v.ID)
	//cache and refresh config by problem-id

	//compile submit
	fileID, result, err := s.middleware.goJudge.Compile([]byte(v.Code), v.Language, uuid.NewString())
	if err != nil {
		submission.Status = u.CompileError
		if result != nil {
			compileError, ok := result.Files["stderr"]
			if !ok {
				submission.CompileMessage = "Compile error"
			} else {
				submission.CompileMessage = string(compileError)
			}
		}
		_, updateErr := submission.Update().Save(ctx)
		marshal, marshalErr := json.Marshal(mq.Ent2mq(submission))
		if marshalErr != nil {
			return marshalErr
		}
		redisErr := s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour).Err()
		if redisErr != nil {
			return err
		}
		if updateErr != nil {
			return updateErr
		}
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
			return err
		}
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
		return err
	}
	//Save into database
	submission.Status = status
	submission.Point = score
	submission.TotalTime = int32(totalTime)
	submission.MaxMemory = int32(maxMemory)
	_, err = submission.Update().Save(ctx)
	marshal, _ := json.Marshal(mq.Ent2mq(submission))
	s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour)
	if err != nil {
		return err
	}
	for _, build := range builders {
		build.SetSubmissionID(submission.ID)
	}
	_, err = s.middleware.ent.SubmissionCase.CreateBulk(builders...).Save(context.Background())
	if err != nil {
		return err
	}
	return nil
}
