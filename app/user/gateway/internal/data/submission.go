package data

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
	"sastoj/app/user/gateway/internal/biz"
)

type submissionRepo struct {
	data *Data
	log  *log.Helper
}

func (s *submissionRepo) CreateSelfTest(ctx context.Context, selfTest *biz.SelfTest) error {
	q, err := s.data.ch.QueueDeclare(
		"self-test", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}
	body, err := json.Marshal(selfTest)
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

func (s *submissionRepo) CreateSubmission(ctx context.Context, submission *biz.Submission) (string, error) {
	// TODO: store into cache
	// store into mq
	q, err := s.data.ch.QueueDeclare(
		"submission", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(submission)
	if err != nil {
		return "", err
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
	s.data.cache.submissions[submission.ID] = submission
	return submission.ID, nil
}

func (s *submissionRepo) GetSubmission(ctx context.Context, submissionID string, userID int64) (*biz.Submission, error) {
	po := s.data.cache.submissions[submissionID]
	if po.UserID != userID {
		return nil, errors.New("permission denied")
	}
	return po, nil
}

// NewSubmissionRepo .
func NewSubmissionRepo(data *Data, logger log.Logger) biz.SubmissionRepo {
	return &submissionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
