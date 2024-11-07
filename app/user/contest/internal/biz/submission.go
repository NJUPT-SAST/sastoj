package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Submission struct {
	ID          string    `json:"id,omitempty"`
	UserID      int64     `json:"user_id,omitempty"`
	ProblemID   int64     `json:"problem_id,omitempty"`
	Code        string    `json:"code,omitempty"`
	Status      int16     `json:"state,omitempty"`
	Point       int16     `json:"point,omitempty"`
	CreateTime  time.Time `json:"create_time"`
	TotalTime   uint64    `json:"total_time,omitempty"`
	MaxMemory   uint64    `json:"max_memory,omitempty"`
	Language    string    `json:"language,omitempty"`
	CaseVersion int8      `json:"case_version,omitempty"`
	CompileMsg  string    `json:"compile_msg,omitempty"`
}

type SelfTest struct {
	ID         string `json:"id,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	ProblemID  int64  `json:"problem_id,omitempty"`
	Code       string `json:"code,omitempty"`
	Language   string `json:"language,omitempty"`
	Input      string `json:"input,omitempty"`
	IsCompiled bool   `json:"is_compiled,omitempty"`
	CompileMsg string `json:"compile_msg,omitempty"`
	Stdout     string `json:"stdout,omitempty"`
	Stderr     string `json:"stderr,omitempty"`
	Time       uint64 `json:"time,omitempty"`
	Memory     uint64 `json:"memory,omitempty"`
}

type Case struct {
	Index  int32  `json:"index,omitempty"`
	Point  int32  `json:"point,omitempty"`
	State  int32  `json:"state,omitempty"`
	Time   uint64 `json:"time,omitempty"`
	Memory uint64 `json:"memory,omitempty"`
}

type Subtask struct {
	Index  int32   `json:"index,omitempty"`
	Point  int32   `json:"point,omitempty"`
	State  int32   `json:"state,omitempty"`
	Time   uint64  `json:"time,omitempty"`
	Memory uint64  `json:"memory,omitempty"`
	Cases  []*Case `json:"cases,omitempty"`
}

type SubmissionRepo interface {
	CreateSubmission(ctx context.Context, submission *Submission) error
	CreateSelfTest(ctx context.Context, pretest *SelfTest) error
	// UpdateSubmission UpdateStatus(ctx context.Context, submissionID string, status int16) error
	UpdateSubmission(ctx context.Context, submission *Submission) error
	GetSubmission(ctx context.Context, submissionID string, contestID int64) (*Submission, error)
	GetSubmissions(ctx context.Context, contestID int64, problemID int64) ([]*Submission, error)
	GetSelfTest(ctx context.Context, submissionID string) (*SelfTest, error)
	GetCases(ctx context.Context, submissionID string, contestID int64) ([]*Subtask, error)
}

type SubmissionUsecase struct {
	repo SubmissionRepo
	log  *log.Helper
}

func NewSubmissionUsecase(repo SubmissionRepo, logger log.Logger) *SubmissionUsecase {
	return &SubmissionUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SubmissionUsecase) CreateSubmission(ctx context.Context, submission *Submission) error {
	return uc.repo.CreateSubmission(ctx, submission)
}

//func (uc *SubmissionUsecase) UpdateStatus(ctx context.Context, submissionID string, status int16) error {
//	return uc.repo.UpdateStatus(ctx, submissionID, status)
//}

func (uc *SubmissionUsecase) UpdateSubmission(ctx context.Context, submission *Submission) error {
	return uc.repo.UpdateSubmission(ctx, submission)
}

func (uc *SubmissionUsecase) GetSubmission(ctx context.Context, submissionID string, contestID int64) (*Submission, error) {
	return uc.repo.GetSubmission(ctx, submissionID, contestID)
}

func (uc *SubmissionUsecase) GetSubmissions(ctx context.Context, contestID int64, problemID int64) ([]*Submission, error) {
	return uc.repo.GetSubmissions(ctx, contestID, problemID)
}

func (uc *SubmissionUsecase) CreateSelfTest(ctx context.Context, selfTest *SelfTest) error {
	return uc.repo.CreateSelfTest(ctx, selfTest)
}

func (uc *SubmissionUsecase) GetSelfTest(ctx context.Context, submissionID string) (*SelfTest, error) {
	return uc.repo.GetSelfTest(ctx, submissionID)
}

func (uc *SubmissionUsecase) GetCases(ctx context.Context, submissionID string, contestID int64) ([]*Subtask, error) {
	return uc.repo.GetCases(ctx, submissionID, contestID)
}
