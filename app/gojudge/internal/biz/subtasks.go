package biz

import (
	"bytes"
	"context"
	"github.com/google/uuid"
	"sastoj/ent"
	"sastoj/ent/problem"
	u "sastoj/pkg/util"
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
	//cache and refresh config by problem-id

	//get problem from ent
	pro, err := s.middleware.ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		First(context.Background())
	if err != nil {
		return err
	}

	//compile submit
	fileID, result, err := s.middleware.goJudge.Compile(v.Code, v.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			err = s.middleware.ent.Submission.Create().
				SetCaseVersion(int8(pro.CaseVersion)).
				SetUsersID(v.UserID).
				SetCode(v.Code).
				SetLanguage(v.Language).
				SetPoint(0).
				SetCompileMessage(string(result.Files["stderr"])).
				SetCreateTime(v.CreateTime).
				SetProblemsID(v.ProblemID).
				Exec(context.Background())
		}
		return err
	}

	//TODO gen case task and score

	var totalTime uint64 = 0
	var maxMemory uint64 = 0
	var score int16 = 0
	var builders []*ent.SubmissionCaseCreate

	for _, task := range s.config.Task.Subtasks {
		//test each case
		var taskBuilder []*ent.SubmissionCaseCreate
		var states []int16
		for index, c := range task.Cases {
			in, ans, err := s.middleware.fileManage.FetchCase(v.ProblemID, c.Input, c.Answer)
			if err != nil {
				_ = s.middleware.goJudge.DeleteFile(fileID)
				return err
			}
			result, err := s.middleware.goJudge.Judge(string(in), v.Language, fileID, uuid.NewString(), uint64(s.config.ResourceLimits.Time), uint64(s.config.ResourceLimits.Time*2), uint64(s.config.ResourceLimits.Memory), int64(len(ans)))
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
				//get problem cases ID
				var problemCasesID int64
				for _, problemCase := range pro.Edges.ProblemCases {
					if problemCase.Index == int16(index) {
						problemCasesID = problemCase.ID
					}
				}
				msg := result.Files["stderr"]
				create := s.middleware.ent.SubmissionCase.Create().
					SetProblemCasesID(problemCasesID).
					SetMemory(int32(result.Memory)).
					SetTime(int32(result.RunTime)).
					SetMessage(string(msg)).
					SetState(state)
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
				if problemCase.Index == int16(index) {
					problemCasesID = problemCase.ID
				}
			}
			//gen builder
			create := s.middleware.ent.SubmissionCase.Create().
				SetProblemCasesID(problemCasesID).
				SetMemory(int32(result.Memory)).
				SetTime(int32(result.RunTime)).
				SetMessage(result.Status.String()).
				SetState(int16(result.Status.Number()))
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
			for _, builder := range taskBuilder {
				builder.SetPoint(score) //TODO check
				builders = append(builders, builder)
			}
		} else {
			for _, builder := range taskBuilder {
				builder.SetPoint(0) //TODO check
				builders = append(builders, builder)
			}
		}
	}

	err = s.middleware.goJudge.DeleteFile(fileID)
	if err != nil {
		return err
	}
	//Save into database

	cases, err := s.middleware.ent.SubmissionCase.CreateBulk(builders...).Save(context.Background())
	if err != nil {
		return err
	}
	var ids []int64
	for _, ca := range cases {
		ids = append(ids, ca.ID)
	}

	err = s.middleware.ent.Submission.Create().
		SetCaseVersion(int8(pro.CaseVersion)).
		SetUsersID(v.UserID).
		SetCode(v.Code).
		SetLanguage(v.Language).
		SetPoint(score).
		SetStatus(1).
		SetCreateTime(v.CreateTime).
		SetTotalTime(int32(totalTime)).
		SetMaxMemory(int32(maxMemory)).
		SetProblemsID(v.ProblemID).
		AddSubmissionCaseIDs(ids...).
		Exec(context.Background())

	if err != nil {
		return err
	}
	return nil
}
