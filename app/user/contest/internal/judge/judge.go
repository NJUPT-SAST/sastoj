package judge

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"sastoj/app/user/contest/internal/biz"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Judge allows all judges to handling messages from an amqp.Channel
type Judge struct {
	log *log.Helper
}

type Submit = biz.Submit
type SelfTest = biz.Pretest

const (
	submit   = "submit"
	selfTest = "selfTest"
)

// StartLoopOnSubmit start a loop for handling messages from an amqp.Channel and return a function to stop the loop.
// The handle function is used to handling Submit, which interact with the judge, then save the result and do not require to return.
func (g Judge) StartLoopOnSubmit(ctx context.Context, ch *amqp.Channel, handle func(ctx context.Context, v *Submit) error) (close func()) {
	control := make(chan bool)
	_, err := ch.QueueDeclare(
		submit,
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
				messages, err := ch.ConsumeWithContext(
					ctx,
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
					if handle(ctx, v) != nil {
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
func (g Judge) StartLoopOnSelfTest(ctx context.Context, ch *amqp.Channel, handle func(ctx context.Context, v *SelfTest) error) (close func()) {
	control := make(chan bool)
	_, err := ch.QueueDeclare(
		selfTest,
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
				messages, err := ch.ConsumeWithContext(
					ctx,
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
					if handle(ctx, v) != nil {
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
