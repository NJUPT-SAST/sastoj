package data

import (
	"context"
	"encoding/json"
	"errors"
	"sastoj/app/user/contest/internal/biz"
	"sastoj/ent/submission"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

type submitRepo struct {
	data *Data
	log  *log.Helper
}

func (s *submitRepo) GetCases(submissionID int64, userID int64) ([]*biz.Case, error) {
	po, err := s.data.redis.Get(context.Background(), "case:"+strconv.FormatInt(userID, 10)+":"+strconv.FormatInt(submissionID, 10)).Result()
	if err != nil {
		return nil, err
	}
	var cases []*biz.Case
	err = json.Unmarshal([]byte(po), &cases)
	if err != nil {
		return nil, err
	}
	return cases, nil
}

func (s *submitRepo) CreatePretest(ctx context.Context, p *biz.Pretest) error {
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

func (s *submitRepo) GetSubmission(ctx context.Context, submitID int64, userID int64) (*biz.Submit, error) {
	po, err := s.data.db.Submission.Get(ctx, submitID)
	if err != nil {
		return nil, err
	}
	if po.UserID != userID {
		// TODO: ADD Global Error: Permission Denied
		return nil, errors.New("permission denied")
	}
	return &biz.Submit{
		ID:          po.ID,
		UserID:      po.UserID,
		ProblemID:   po.ProblemID,
		Code:        po.Code,
		Status:      po.Status,
		Point:       po.Point,
		CreateTime:  po.CreateTime,
		TotalTime:   po.TotalTime,
		MaxMemory:   po.MaxMemory,
		Language:    po.Language,
		CaseVersion: po.CaseVersion,
	}, nil
}

func (s *submitRepo) GetSubmissions(ctx context.Context, userId int64, problemId int64) ([]*biz.Submit, error) {
	submissions, err := s.data.db.Submission.Query().
		Select(submission.FieldID, submission.FieldStatus, submission.FieldPoint, submission.FieldLanguage, submission.FieldCreateTime).
		Where(submission.UserIDEQ(userId), submission.ProblemIDEQ(problemId)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var submits []*biz.Submit
	for _, v := range submissions {
		submits = append(submits, &biz.Submit{
			ID:         v.ID,
			UserID:     userId,
			ProblemID:  problemId,
			Status:     v.Status,
			Point:      v.Point,
			CreateTime: v.CreateTime,
		})
	}
	return submits, nil
}

func (s *submitRepo) CreateSubmit(ctx context.Context, submit *biz.Submit) (int64, error) {
	po, err := s.data.db.Submission.Create().
		SetUsersID(submit.UserID).
		SetProblemsID(submit.ProblemID).
		SetCode(submit.Code).
		SetStatus(submit.Status).
		SetPoint(submit.Point).
		SetCreateTime(submit.CreateTime).
		SetTotalTime(submit.TotalTime).
		SetMaxMemory(submit.MaxMemory).
		SetLanguage(submit.Language).
		SetCaseVersion(submit.CaseVersion).
		Save(ctx)
	if err != nil {
		return 0, err
	}
	q, err := s.data.ch.QueueDeclare(
		"submission", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
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

func (s *submitRepo) UpdateStatus(ctx context.Context, submitID int64, status int16) error {
	_, err := s.data.db.Submission.UpdateOneID(submitID).
		SetStatus(status).
		Save(ctx)
	return err
}

func (s *submitRepo) UpdateSubmit(ctx context.Context, submit *biz.Submit) error {
	_, err := s.data.db.Submission.UpdateOneID(submit.ID).
		SetStatus(submit.Status).
		SetPoint(submit.Point).
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
