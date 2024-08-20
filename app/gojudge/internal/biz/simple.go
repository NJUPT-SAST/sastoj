package biz

import (
	"bytes"
	"context"
	"errors"
	"sastoj/ent"
	"sastoj/ent/problem"
	"sastoj/pkg/mq"
	u "sastoj/pkg/util"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Simple Adapter
type Simple struct {
	middleware *Middleware
	config     *u.JudgeConfig
}

// handleSubmit process the submission:
// 1. get problem, problem-cases and config-file and by problem-id
// 2. compile submission
// 2. get in-out by problem-id and case-id
// 3. judge each case-submission and save, then create submission and save
func (s *Simple) handleSubmit(v *Submit) (*mq.Submission, []*ent.SubmissionCaseCreate, error) {
	ctx := context.Background()

	submission := &mq.Submission{
		ID:         v.ID,
		UserID:     v.UserID,
		ProblemID:  v.ProblemID,
		Code:       v.Code,
		Status:     u.Judging,
		Point:      0,
		CreateTime: v.CreateTime,
		TotalTime:  0,
		MaxMemory:  0,
		Language:   v.Language,
		Stderr:     "",
		Token:      v.Token,
	}

	//get problem from ent
	pro, err := s.middleware.ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		WithProblemCases().
		First(ctx)
	if err != nil {
		submission.Status = u.SystemError
		return submission, nil, errors.New("problem not found")
	}

	submission.CaseVer = pro.CaseVersion

	// TODO: cache and refresh config by problem-id

	//compile submit
	fileID, result, err := s.middleware.goJudge.Compile([]byte(v.Code), v.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			submission.Status = u.CompileError
			submission.Stderr = "Compile error without stderr"
			compileError, ok := result.Files["stderr"]
			if ok {
				submission.Stderr = string(compileError)
			}
		} else {
			submission.Status = u.SystemError
		}
		return submission, nil, err
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
			submission.Status = u.Invalid
			_ = s.middleware.goJudge.DeleteFile(fileID)
			return submission, nil, err
		}

		in, ans, err := s.middleware.fileManage.FetchCase(v.ProblemID, c.Input, c.Answer)
		//TODO remove
		in = u.RemoveCr(in)
		ans = u.RemoveCr(ans)
		if err != nil {
			submission.Status = u.SystemError
			_ = s.middleware.goJudge.DeleteFile(fileID)
			return submission, nil, err
		}

		result, err := s.middleware.goJudge.Judge(in, v.Language, fileID, uuid.NewString(), uint64(s.config.ResourceLimits.Time), uint64(s.config.ResourceLimits.Time*2), uint64(s.config.ResourceLimits.Memory), int64(len(ans)))
		if err != nil {
			submission.Status = u.RuntimeError
			submission.Stderr = err.Error()
			gojudgeErr := s.middleware.goJudge.DeleteFile(fileID)
			if gojudgeErr != nil {
				s.middleware.logger.Errorf("delete gojudge file error: %v", gojudgeErr)
			}
			return submission, nil, err
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
		if problemCasesID == 0 {
			submission.Status = u.SystemError
			_ = s.middleware.goJudge.DeleteFile(fileID)
			return submission, nil, errors.New("problem case not found")
		}

		//gen builder
		create := s.middleware.ent.SubmissionCase.Create().
			SetProblemCasesID(problemCasesID).
			SetMemory(int32(result.Memory)).
			SetTime(int32(result.RunTime)).
			SetMessage(result.Status.String()).
			SetState(state).
			SetPoint(c.Score)
		builders = append(builders, create)
	}

	err = s.middleware.goJudge.DeleteFile(fileID)
	if err != nil {
		s.middleware.logger.Errorf("delete gojudge file error: %v", err)
	}

	submission.Status = status
	submission.Point = score
	submission.TotalTime = totalTime
	submission.MaxMemory = maxMemory

	return submission, builders, nil
}
