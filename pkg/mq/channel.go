package mq

import (
	"context"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Channel interface {
	Publish(ctx context.Context, v interface{}) error
	Close() error
}

type OjChannel struct {
	ch *amqp.Channel
	q  *amqp.Queue
}

type OjExchange struct {
	ch    *amqp.Channel
	topic string
}

func NewChannel(ch *amqp.Channel, name string) (*OjChannel, error) {
	q, err := ch.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}
	return &OjChannel{
		ch: ch,
		q:  &q,
	}, nil
}

func NewExchange(ch *amqp.Channel, name string) (*OjExchange, error) {
	err := ch.ExchangeDeclare(
		name,    // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return nil, err
	}
	return &OjExchange{
		ch:    ch,
		topic: name,
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

func (o *OjExchange) Publish(ctx context.Context, v any) error {
	body, err := json.Marshal(v)
	if err != nil {
		return err
	}
	log.Infof("publish message: %s", body)
	return o.ch.PublishWithContext(ctx, o.topic, o.topic, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}

func (o *OjExchange) Close() error {
	return o.ch.Close()
}
