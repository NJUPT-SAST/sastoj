package biz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sastoj/ent"
	"sastoj/ent/problem"
	u "sastoj/pkg/util"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

// Subtasks Adapter
type Subtasks struct {
	middleware *Middleware
	config     *u.JudgeConfig
}

// handleSubmit process the submission
// 1. get problem, problem-cases and config-file and by problem-id
// 2. compile submission
// 2. get in-out by problem-id and case-id
// 3. judge each case-submission and save, then create submission and save
func (s *Subtasks) handleSubmit(v *Submit) error {
	ctx := context.Background()

	//get problem from ent
	pro, err := s.middleware.ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		WithProblemCases().
		First(ctx)
	if err != nil {
		return err
	}

	// TODO: cache and refresh config by problem-id

	redisKey := fmt.Sprintf("submission:%d:%s", v.UserID, v.ID)

	//compile submit
	fileID, result, err := s.middleware.goJudge.Compile([]byte(v.Code), v.Language, uuid.NewString())
	if err != nil {
		if result != nil {
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
			v.ID = strconv.FormatInt(save.ID, 10)
			v.Status = u.CompileError
			marshal, marshalErr := json.Marshal(v)
			if marshalErr != nil {
				log.Errorf("marshal error: %v", marshalErr)
			}
			redisErr := s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour).Err()
			if redisErr != nil {
				log.Errorf("redis error: %v", redisErr)
			}
			return err
		}
	}

	var totalTime uint64 = 0
	var maxMemory uint64 = 0
	var score int16 = 0
	var status int16 = u.Accepted
	var builders = make([]*ent.SubmissionCaseCreate, 0)

	for _, task := range s.config.Task.Subtasks {
		//test each case
		var taskBuilder []*ent.SubmissionCaseCreate
		var states []int16
		for _, c := range task.Cases {

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
				taskBuilder = append(taskBuilder, create)
				states = append(states, state)
				continue
			}
			//Accept
			totalTime += result.RunTime
			if maxMemory < result.Memory {
				maxMemory = result.Memory
			}
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
				SetPoint(c.Score).
				SetState(state)
			taskBuilder = append(taskBuilder, create)
			states = append(states, state)
		}
		var res int16 = u.Accepted
		for _, state := range states {
			if state != u.Accepted {
				res = state
			}
		}
		if res == u.Accepted {
			score += task.Score
			builders = append(builders, taskBuilder...)
		} else {
			for _, builder := range taskBuilder {
				builder.SetPoint(0)
			}
			builders = append(builders, taskBuilder...)
			continue
			//TODO if oi break
		}
	}

	err = s.middleware.goJudge.DeleteFile(fileID)
	if err != nil {
		return err
	}

	//Save into database
	submission, err := s.middleware.ent.Submission.Create().
		SetCaseVersion(int8(pro.CaseVersion)).
		SetUsersID(v.UserID).
		SetCode(v.Code).
		SetLanguage(v.Language).
		SetPoint(score).
		SetStatus(status).
		SetCreateTime(v.CreateTime).
		SetTotalTime(int32(totalTime)).
		SetMaxMemory(int32(maxMemory)).
		SetCompileMessage("").
		SetProblemsID(v.ProblemID).
		Save(ctx)
	if err != nil {
		return err
	}

	v.ID = strconv.FormatInt(submission.ID, 10)
	v.Status = status
	v.Point = score
	v.TotalTime = int32(totalTime)
	v.MaxMemory = int32(maxMemory)
	marshal, err := json.Marshal(v)
	if err != nil {
		log.Errorf("marshal error: %v", err)
	}
	s.middleware.redis.Set(ctx, redisKey, marshal, 2*time.Hour)

	for _, build := range builders {
		build.SetSubmissionID(submission.ID)
	}
	_, err = s.middleware.ent.SubmissionCase.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
