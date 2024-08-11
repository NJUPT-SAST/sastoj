package data

import (
	"context"
	"encoding/json"
	"errors"
	pb "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type submissionDTO struct {
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

type selfTestDTO struct {
	ID       string `json:"id,omitempty"`
	UserID   int64  `json:"user_id,omitempty"`
	Code     string `json:"code,omitempty"`
	Language string `json:"language,omitempty"`
	Input    string `json:"input,omitempty"`
	Token    string `json:"secret,omitempty"`
}

type submissionRepo struct {
	data *Data
	log  *log.Helper
}

func (s *submissionRepo) GetCases(ctx context.Context, userID int64, submissionID string) ([]*biz.Case, error) {
	cases, ok := s.data.cache.cases[submissionID]
	if !ok {
		rev, err := s.data.cc.GetCases(ctx, &pb.GetCasesRequest{
			SubmissionId: submissionID,
		})

		if err != nil {
			return nil, err
		}
		for _, v := range rev.Cases {
			cases = append(cases, &biz.Case{
				Index: v.Index,
				State: v.State,
			})
		}
	}
	return cases, nil
}

func (s *submissionRepo) UpdateSubmission(ctx context.Context, submission *biz.Submission) error {
	s.data.cache.submissions[submission.ID].Point = submission.Point
	s.data.cache.submissions[submission.ID].Status = submission.Status
	s.data.cache.submissions[submission.ID].TotalTime = submission.TotalTime
	s.data.cache.submissions[submission.ID].MaxMemory = submission.MaxMemory
	return nil
}

func (s *submissionRepo) UpdateSelfTest(ctx context.Context, selfTest *biz.SelfTest) error {
	s.data.cache.selfTests[selfTest.ID].Output = selfTest.Output
	return nil
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
	sDTO := &selfTestDTO{
		ID:       selfTest.ID,
		UserID:   selfTest.UserID,
		Code:     selfTest.Code,
		Language: selfTest.Language,
		Input:    selfTest.Input,
		Token:    s.data.cache.token,
	}
	body, err := json.Marshal(sDTO)
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
	if err != nil {
		return err
	}
	s.data.cache.selfTests[selfTest.ID] = selfTest
	return nil
}

func (s *submissionRepo) CreateSubmission(ctx context.Context, submission *biz.Submission) (string, error) {
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
	sDTO := &submissionDTO{
		ID:         submission.ID,
		UserID:     submission.UserID,
		ProblemID:  submission.ProblemID,
		Code:       submission.Code,
		Status:     submission.Status,
		Point:      submission.Point,
		CreateTime: submission.CreateTime,
		TotalTime:  submission.TotalTime,
		MaxMemory:  submission.MaxMemory,
		Language:   submission.Language,
		Token:      s.data.cache.token,
	}
	body, err := json.Marshal(sDTO)
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
