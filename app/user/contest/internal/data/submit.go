package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"sastoj/app/user/contest/internal/biz"
)

type submitRepo struct {
	data *Data
	log  *log.Helper
}

func (s submitRepo) CreatePretest(ctx context.Context, p *biz.Pretest) error {
	q, err := s.data.ch.QueueDeclare(
		"pretest", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}
	body, err := json.Marshal(p)
	if err != nil {
		return err
	}
	err = s.data.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	return err
}

func (s submitRepo) GetSubmission(ctx context.Context, submitID int64, userID int64) (*biz.Submit, error) {
	po, err := s.data.db.Submit.Get(ctx, submitID)
	if err != nil {
		return nil, err
	}
	if po.UserID != userID {
		return nil, errors.New("permission denied")
	}
	return &biz.Submit{
		ID:          po.ID,
		UserID:      po.UserID,
		ProblemID:   po.ProblemID,
		Code:        po.Code,
		Status:      int32(po.Status),
		Point:       int32(po.Point),
		CreateTime:  po.CreateTime,
		TotalTime:   po.TotalTime,
		MaxMemory:   po.MaxMemory,
		Language:    po.Language,
		CaseVersion: int32(po.CaseVersion),
	}, nil
}

func (s submitRepo) CreateSubmit(ctx context.Context, submit *biz.Submit) (int64, error) {
	po, err := s.data.db.Submit.Create().
		SetUsersID(submit.UserID).
		SetProblemsID(submit.ProblemID).
		SetCode(submit.Code).
		SetStatus(int16(submit.Status)).
		SetPoint(int16(submit.Point)).
		SetCreateTime(submit.CreateTime).
		SetTotalTime(submit.TotalTime).
		SetMaxMemory(submit.MaxMemory).
		SetLanguage(submit.Language).
		SetCaseVersion(int8(submit.CaseVersion)).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	q, err := s.data.ch.QueueDeclare(
		"submit", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return 0, err
	}
	body, err := json.Marshal(po)
	if err != nil {
		return 0, err
	}
	err = s.data.ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	return po.ID, nil
}

func (s submitRepo) UpdateStatus(ctx context.Context, submitID int64, status int32) error {
	_, err := s.data.db.Submit.UpdateOneID(submitID).
		SetStatus(int16(status)).
		Save(ctx)
	return err
}

func (s submitRepo) UpdateSubmit(ctx context.Context, submit *biz.Submit) error {
	_, err := s.data.db.Submit.UpdateOneID(submit.ID).
		SetStatus(int16(submit.Status)).
		SetPoint(int16(submit.Point)).
		SetTotalTime(submit.TotalTime).
		SetMaxMemory(submit.MaxMemory).
		Save(ctx)
	return err
}

func NewSubmitRepo(data *Data, logger log.Logger) biz.SubmitRepo {
	return &submitRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
