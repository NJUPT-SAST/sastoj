package service

import (
	"context"
	"sastoj/app/judge/freshcup/internal/biz"
	"sastoj/pkg/mq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
)

// SubmissionService is a submission service.
type SubmissionService struct {
	uc *biz.SubmissionUsecase
}

// NewSubmissionService new a submission service.
func NewSubmissionService(uc *biz.SubmissionUsecase) *SubmissionService {
	return &SubmissionService{uc: uc}
}

func (s *SubmissionService) SubmissionHandle(ctx context.Context, topic string, headers broker.Headers, msg *mq.Submission) error {
	log.Infof("Topic %s, Headers: %+v, Payload: %+v\n", topic, headers, msg)
	err := s.uc.JudgeSubmission(ctx, msg)
	if err != nil {
		log.Errorf("JudgeSubmission error: %v", err)
	}
	return err
}

func (s *SubmissionService) SelfTestHandle(ctx context.Context, topic string, headers broker.Headers, msg *mq.SelfTest) error {
	log.Infof("Topic %s, Headers: %+v, Payload: %+v\n", topic, headers, msg)
	err := s.uc.JudgeSelfTest(ctx, msg)
	if err != nil {
		log.Errorf("JudgeSelfTest error: %v", err)
	}
	return err
}
