package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"sastoj/app/gojudge/internal/data"
	"sastoj/ent"
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
			sum := config.Score
			var num int16 = 0
			for _, c := range config.Task.Cases {
				if c.Score == 0 {
					num += 1
				} else {
					sum -= c.Score
				}
			}
			if sum > 0 {
				avg := sum / num
				for i := range config.Task.Cases {
					c := &config.Task.Cases[i]
					if c.Score == 0 {
						c.Score = avg
					}
				}
			}
			simple := Simple{
				middleware: m,
				config:     config,
			}
			return simple.handleSubmit(v)
		case "subtasks":
			sum := config.Score
			var num int16 = 0
			for _, sub := range config.Task.Subtasks {
				if sub.Score == 0 {
					num += 1
				} else {
					sum -= sub.Score
				}
			}
			if sum > 0 {
				avg := sum / num
				for i := range config.Task.Subtasks {
					sub := &config.Task.Subtasks[i]
					if sub.Score == 0 {
						sub.Score = avg
					}
					subAvg := sub.Score / int16(len(sub.Cases))
					for j := range sub.Cases {
						c := &sub.Cases[j]
						if c.Score == 0 {
							c.Score = subAvg
						}
					}
				}
			}
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
	commands := *m.goJudge.commands
	testConfig := commands["test"]
	fileID, result, err := m.goJudge.Compile(v.Code, v.Language, uuid.NewString())
	if err != nil {
		if result != nil {
			//TODO result struct
			m.redis.Set(context.Background(), v.ID, result, 10000)
		}
		return err
	}
	result, err = m.goJudge.Judge(v.Input, v.Language, fileID, uuid.NewString(), testConfig.CompileConfig.CpuTimeLimit, testConfig.CompileConfig.ClockTimeLimit, testConfig.CompileConfig.MemoryLimit, testConfig.CompileConfig.StdoutMaxSize)
	if err != nil {
		_ = m.goJudge.DeleteFile(fileID)
		return err
	}
	err = m.goJudge.DeleteFile(fileID)
	if err != nil {
		return err
	}
	//TODO result struct
	m.redis.Set(context.Background(), v.ID, result, 10000)
	return nil
}