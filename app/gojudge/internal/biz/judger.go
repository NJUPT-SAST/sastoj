package biz

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Judge allows all judges to handling messages from an amqp.Channel
type Judge struct {
	log *log.Helper
}

type Submit struct {
	ID         string    `json:"id,omitempty"`
	UserID     int64     `json:"user_id,omitempty"`
	ProblemID  int64     `json:"problem_id,omitempty"`
	Code       string    `json:"code,omitempty"`
	Status     int16     `json:"state,omitempty"`
	Point      int16     `json:"point,omitempty"`
	CreateTime time.Time `json:"create_time"`
	TotalTime  int32     `json:"total_time,omitempty"`
	MaxMemory  int32     `json:"max_memory,omitempty"`
	Language   string    `json:"language,omitempty"`
	Token      string    `json:"secret,omitempty"`
}

type SelfTest struct {
	ID       string `json:"id,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Language string `json:"language,omitempty"`
	Input    string `json:"input,omitempty"`
	Token    string `json:"secret,omitempty"`
}

// StartLoopOnSubmit start a loop for handling messages from an amqp.Channel and return a function to stop the loop.
// The handle function is used to handling Submit, which interact with the judge, then save the result and do not require to return.
func (g *Judge) StartLoopOnSubmit(ch *amqp.Channel, topic string, handle func(v *Submit) error) (close func()) {
	control := make(chan bool)
	_, err := ch.QueueDeclare(
		topic,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		g.log.Errorf("Go-Judge Consumer(OnSubmit) Start Failed: Failed to declare a queue: %s", err)
	}
	go func() {
		for {
			select {
			case <-control:
				return
			default:
				messages, err := ch.Consume(
					submit,
					"",
					true,
					false,
					false,
					false,
					nil,
				)
				if err != nil {
					g.log.Errorf("Go-Judge Consumer(OnSubmit) Error: Failed consume message: %s", err)
				}
				for d := range messages {
					body := d.Body
					g.log.Infof("Go-Judge Consumer(OnSubmit) Received a message: %s", body)
					v := &Submit{}
					err := json.Unmarshal(body, v)
					if err != nil {
						g.log.Errorf("Go-Judge Consumer(OnSubmit) Error: Failed to de-serialise a submit: %s", err)
					}
					if handle(v) != nil {
						g.log.Errorf("Go-Judge Consumer(OnSubmit) Error: Failed to handle a submit: %s", err)
					}
				}
			}
		}
	}()
	return func() {
		control <- false
	}
}

// StartLoopOnSelfTest start a loop for handling messages from an amqp.Channel and return a function to stop the loop.
// The handle function is used to handling SelfTest, which interact with the judge, then save the result and do not require to return.
func (g *Judge) StartLoopOnSelfTest(ch *amqp.Channel, topic string, handle func(v *SelfTest) error) (close func()) {
	control := make(chan bool)
	_, err := ch.QueueDeclare(
		topic,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		g.log.Errorf("Go-Judge Consumer(OnSelfTest) Start Failed: Failed to declare a queue: %s", err)
	}
	go func() {
		for {
			select {
			case <-control:
				return
			default:
				messages, err := ch.Consume(
					selfTest,
					"",
					true,
					false,
					false,
					false,
					nil,
				)
				if err != nil {
					g.log.Errorf("Go-Judge Consumer(OnSelfTest) Error: Failed consume message: %s", err)
				}
				for d := range messages {
					body := d.Body
					g.log.Infof("Go-Judge Consumer(OnSelfTest) Received a message: %s", body)
					v := &SelfTest{}
					err := json.Unmarshal(body, v)
					if err != nil {
						g.log.Errorf("Go-Judge Consumer(OnSelfTest) Error: Failed to de-serialise a self test: %s", err)
					}
					if handle(v) != nil {
						g.log.Errorf("Go-Judge Consumer(OnSelfTest) Error: Failed to handle a self test: %s", err)
					}
				}
			}
		}
	}()
	return func() {
		control <- false
	}
}
