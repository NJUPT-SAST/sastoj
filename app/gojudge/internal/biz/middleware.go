package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"sastoj/app/gojudge/internal/data"
	"sastoj/ent"
	"sastoj/ent/problem"
	"slices"
)

const (
	// submit topic name
	submit = "submission"
	// selfTest topic name
	selfTest = "self-test"
)

// Middleware is a Client for a single go-judge
type Middleware struct {
	Ch         *amqp.Channel
	Ent        *ent.Client
	FileManage *data.FileManage
	goJudge    *GoJudge
	Logger     *log.Helper
	close      []func()
}

func NewMiddleware(data *data.Data) *Middleware {
	return &Middleware{
		Ch:         data.Ch,
		Ent:        data.Ent,
		FileManage: data.FileManage,
		goJudge: &GoJudge{
			client:   data.Client,
			commands: data.Commands,
		},
		Logger: data.Logger,
		close:  []func(){},
	}
}

func (m *Middleware) Start() {
	m.StartLoopOnSubmit()
	m.StartLoopOnSelfTest()
}

func (m *Middleware) StartLoopOnSubmit() {
	c := StartLoopOnSubmit(m.Logger, m.Ch, submit, m.handleSubmit)
	m.close = append(m.close, c)
}

func (m *Middleware) StartLoopOnSelfTest() {
	c := StartLoopOnSelfTest(m.Logger, m.Ch, selfTest, m.handleSelfTest)
	m.close = append(m.close, c)
}

func (m *Middleware) Close() {
	for _, c := range m.close {
		c()
	}
}

// handleSubmit process the submission
// 1. get problem, problem-cases and config-file and by problem-id
// 2. compile submission
// 2. get in-out by problem-id and case-id
// 3. judge each case-submission and save, then create submission and save
func (m *Middleware) handleSubmit(v *Submit) error {

	//cache and refresh config by problem-id
	config, err := m.FileManage.GetConfig(v.ProblemID)
	if err != nil {
		return err
	}
	//get problem from ent
	pro, err := m.Ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		First(context.Background())
	if err != nil {
		return err
	}
	//compile submit
	fileID, result, err := m.goJudge.Compile(v.Code, v.Language, "")
	if err != nil {
		//compile err
		return err
	}
	//TODO 处理编译错误
	fmt.Print(result)

	totalTime := uint64(0)
	maxMemory := uint64(0)
	score := 0
	var builder []*ent.SubmissionCaseCreate

	//test each case
	for index, c := range config.Task.Cases {
		in, ans, err := m.FileManage.FetchCase(v.ProblemID, c.Input, c.Answer)
		if err != nil {
			_ = m.goJudge.DeleteFile(fileID)
			return err
		}
		result, err := m.goJudge.Judge(string(in), v.Language, fileID, "", uint64(config.ResourceLimits.Time), uint64(config.ResourceLimits.Time*2), uint64(config.ResourceLimits.Memory), int64(len(ans)))
		if err != nil {
			_ = m.goJudge.DeleteFile(fileID)
			return err
		}
		//non-Accept
		if result.Status.Number() != 1 || !slices.Equal(result.Files["stdout"], ans) {
			//get problem cases ID
			var problemCasesID int64
			for _, problemCase := range pro.Edges.ProblemCases {
				if problemCase.Index == int16(index) {
					problemCasesID = problemCase.ID
				}
			}
			create := m.Ent.SubmissionCase.Create().
				SetProblemCasesID(problemCasesID).
				SetMemory(int32(result.Memory)).
				SetTime(int32(result.RunTime)).
				SetMessage(result.Status.String()).
				SetState(int16(result.Status.Number())).
				SetPoint(0)
			builder = append(builder, create)
			break
		}

		//Accept
		totalTime += result.RunTime
		if maxMemory < result.Memory {
			maxMemory = result.Memory
		}
		score += int(v.Point)

		//get problem cases ID
		var problemCasesID int64
		for _, problemCase := range pro.Edges.ProblemCases {
			if problemCase.Index == int16(index) {
				problemCasesID = problemCase.ID
			}
		}

		//gen builder
		create := m.Ent.SubmissionCase.Create().
			SetProblemCasesID(problemCasesID).
			SetMemory(int32(result.Memory)).
			SetTime(int32(result.RunTime)).
			SetMessage(result.Status.String()).
			SetState(int16(result.Status.Number())).
			SetPoint(v.Point)
		builder = append(builder, create)
	}
	err = m.goJudge.DeleteFile(fileID)
	if err != nil {
		return err
	}
	//Save into database

	cases, err := m.Ent.SubmissionCase.CreateBulk(builder...).Save(context.Background())
	if err != nil {
		return err
	}
	var ids []int64
	for _, ca := range cases {
		ids = append(ids, ca.ID)
	}

	err = m.Ent.Submission.Create().
		SetCaseVersion(int8(pro.CaseVersion)).
		SetUsersID(v.UserID).
		SetCode(v.Code).
		SetLanguage(v.Language).
		SetPoint(int16(score)).
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

// handleSubmit process the self-test
// 1. compile submission
// 2. judge each case-submission and save as self-test
func (m *Middleware) handleSelfTest(v *SelfTest) error {
	//TODO self-test
	return nil
}
