package mq

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	SubmissionQueue = "submission"
	SelfTestQueue   = "self-test"
)

type Channel interface {
	Publish(ctx context.Context, v interface{}) error
	Close() error
}

type OjChannel struct {
	ch *amqp.Channel
	q  *amqp.Queue
}

func newChannel(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func NewSubmissionChannel(ch *amqp.Channel) (*OjChannel, error) {
	q, err := newChannel(ch, SubmissionQueue)
	if err != nil {
		return nil, err
	}
	return &OjChannel{
		ch: ch,
		q:  &q,
	}, nil
}

func NewSelfTestChannel(ch *amqp.Channel) (*OjChannel, error) {
	q, err := newChannel(ch, SelfTestQueue)
	if err != nil {
		return nil, err
	}
	return &OjChannel{
		ch: ch,
		q:  &q,
	}, nil
}

func (o *OjChannel) Publish(ctx context.Context, v any) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	log.Infof("publish message: %s", body)
	return o.ch.PublishWithContext(ctx, "", o.q.Name, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

func (o *OjChannel) Close() error {
	return o.ch.Close()
}
