package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sastoj/app/gojudge/internal/data"
	"sastoj/ent"
	"sastoj/pkg/mq"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

const (
	// submit topic name
	submit = "submission"
	// selfTest topic name
	selfTest = "self-test"
)

type Judge interface {
	handleSubmit(v *Submit) error
	handleSelfTest(v *SelfTest) error
}

// Middleware is a Client for a single go-judge
type Middleware struct {
	ch         *amqp.Channel
	ent        *ent.Client
	redis      *redis.Client
	fileManage *data.FileManage
	goJudge    *GoJudge
	logger     *log.Helper
	close      []func()
}

func NewMiddleware(data *data.Data) *Middleware {
	return &Middleware{
		ch:         data.Ch,
		ent:        data.Ent,
		fileManage: data.FileManage,
		redis:      data.Redis,
		goJudge: &GoJudge{
			client:   data.Client,
			commands: data.Commands,
		},
		logger: log.NewHelper(log.With(log.GetLogger(), "module", "judge-middleware")),
		close:  []func(){},
	}
}

func (m *Middleware) Start() {
	m.StartLoopOnSubmit()
	m.StartLoopOnSelfTest()
}

func (m *Middleware) StartLoopOnSubmit() {
	c := StartLoopOnSubmit(m.logger, m.ch, submit, m.handleSubmit)
	m.close = append(m.close, c)
}

func (m *Middleware) StartLoopOnSelfTest() {
	c := StartLoopOnSelfTest(m.logger, m.ch, selfTest, m.handleSelfTest)
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
	config, err := m.fileManage.GetConfig(v.ProblemID)
	if err != nil {
		return err
	}
	switch config.Judge.JudgeType {
	case "classic":
		switch config.Task.TaskType {
		case "simple":
			simple := Simple{
				middleware: m,
				config:     config,
			}
			submission, builders, err := simple.handleSubmit(v)
			save, entErr := m.ent.Submission.Create().
				SetUserID(v.UserID).
				SetProblemID(v.ProblemID).
				SetCode(v.Code).
				SetStatus(submission.Status).
				SetPoint(submission.Point).
				SetCreateTime(submission.CreateTime).
				SetTotalTime(int32(submission.TotalTime)).
				SetMaxMemory(int32(submission.MaxMemory)).
				SetLanguage(v.Language).
				SetCompileMessage(submission.Stderr).
				SetCaseVersion(int8(submission.CaseVer)).
				Save(context.Background())
			if entErr != nil {
				m.logger.Errorf("save error: %v", err)
			} else {
				submission.ID = strconv.FormatInt(save.ID, 10)
			}
			marshal, _ := json.Marshal(submission)
			redisKey := fmt.Sprintf("submission:%d:%s", v.UserID, v.ID)
			redisErr := m.redis.Set(context.Background(), redisKey, marshal, 30*time.Minute).Err()
			if redisErr != nil {
				m.logger.Errorf("redis error: %v", redisErr)
			}
			if err != nil {
				return err
			}
			for _, builder := range builders {
				builder.SetSubmissionID(save.ID)
			}
			_, err = m.ent.SubmissionCase.CreateBulk(builders...).Save(context.Background())
			if err != nil {
				return err
			}

		case "subtasks":
			subtasks := Subtasks{
				middleware: m,
				config:     config,
			}
			return subtasks.handleSubmit(v)
		}
	}
	return errors.New("invalid judge type")
}

// handleSubmit process the self-test
// 1. compile submission
// 2. judge each case-submission and save as self-test
func (m *Middleware) handleSelfTest(v *SelfTest) error {
	ctx := context.Background()
	commands := *m.goJudge.commands
	testConfig, ok := commands[v.Language]
	if !ok {
		sT := &mq.SelfTest{
			ID:         v.ID,
			UserID:     v.UserID,
			Code:       v.Code,
			Language:   v.Language,
			Input:      v.Input,
			IsCompiled: false,
			Stdout:     "",
			Stderr:     "language not supported",
			Time:       0,
			Memory:     0,
			Token:      "",
		}
		marshal, err := json.Marshal(sT)
		if err != nil {
			m.logger.Errorf("marshal error: %v", err)
		}
		m.redis.Set(ctx, fmt.Sprintf("self-test:%d:%s", v.UserID, v.ID), marshal, 30*time.Minute)
		return errors.New("language not supported")
	}
	fileID, result, err := m.goJudge.Compile([]byte(v.Code), v.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			res, ok := result.Files["stderr"]
			if ok {
				v.Output = string(res)
			} else {
				v.Output = "compile error"
			}
			sT := &mq.SelfTest{
				ID:         v.ID,
				UserID:     v.UserID,
				Code:       v.Code,
				Language:   v.Language,
				Input:      v.Input,
				IsCompiled: false,
				Stdout:     "",
				Stderr:     v.Output,
				Time:       0,
				Memory:     0,
				Token:      "",
			}
			marshal, err := json.Marshal(sT)
			if err != nil {
				m.logger.Errorf("marshal error: %v", err)
			}
			m.redis.Set(ctx, fmt.Sprintf("self-test:%d:%s", v.UserID, v.ID), marshal, 30*time.Minute)
		}
		return err
	}
	result, err = m.goJudge.Judge([]byte(v.Input), v.Language, fileID, uuid.NewString(), testConfig.CompileConfig.CpuTimeLimit, testConfig.CompileConfig.ClockTimeLimit, testConfig.CompileConfig.MemoryLimit, testConfig.CompileConfig.StdoutMaxSize)
	if err != nil {
		_ = m.goJudge.DeleteFile(fileID)
		return err
	}
	err = m.goJudge.DeleteFile(fileID)
	if err != nil {
		return err
	}
	t := result.Time
	memory := result.Memory
	stdout := string(result.Files["stdout"])
	stderr := string(result.Files["stderr"])
	sT := &mq.SelfTest{
		ID:         v.ID,
		UserID:     v.UserID,
		Code:       v.Code,
		Language:   v.Language,
		Input:      v.Input,
		IsCompiled: true,
		Stdout:     stdout,
		Stderr:     stderr,
		Time:       t,
		Memory:     memory,
		Token:      "",
	}
	marshal, err := json.Marshal(sT)
	if err != nil {
		m.logger.Errorf("marshal error: %v", err)
	}
	m.redis.Set(ctx, fmt.Sprintf("self-test:%d:%s", v.UserID, v.ID), marshal, 30*time.Minute)
	return nil
}
