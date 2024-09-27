package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mime/multipart"
)

type Case struct {
	Id        int64
	ProblemId int64
	Point     int32
	Index     int32
	IsAuto    bool
}

type CaseRepo interface {
	UploadCasesFile(int64, multipart.File, string, string) error
}

type CaseUsecase struct {
	repo CaseRepo
	log  *log.Helper
}

func NewCaseUsecase(repo CaseRepo, logger log.Logger) *CaseUsecase {
	return &CaseUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CaseUsecase) DeleteCasesByProblemId(ctx context.Context, pi int64) error {
	uc.log.WithContext(ctx).Infof("DeleteCasesByProblemId: %v", pi)
	return nil
}

func (uc *CaseUsecase) UploadCases(ctx context.Context, problemId int64, casesFile multipart.File, filename string, casesType string) error {
	uc.log.WithContext(ctx).Infof("start uploading %v from %v", filename, casesType)
	err := uc.repo.UploadCasesFile(problemId, casesFile, filename, casesType)
	if err != nil {
		return err
	}
	uc.log.WithContext(ctx).Infof("Upload %v completed.", filename)
	return nil
}
