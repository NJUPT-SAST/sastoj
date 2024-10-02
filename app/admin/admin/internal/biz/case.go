package biz

import (
	"context"
	"mime/multipart"

	"github.com/go-kratos/kratos/v2/log"
)

type Case struct {
	ID        int64
	ProblemID int64
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

func (uc *CaseUsecase) DeleteCasesByProblemID(ctx context.Context, pi int64) error {
	uc.log.WithContext(ctx).Infof("DeleteCasesByProblemID: %v", pi)
	return nil
}

func (uc *CaseUsecase) UploadCases(ctx context.Context, problemID int64, casesFile multipart.File, filename string, casesType string) error {
	uc.log.WithContext(ctx).Infof("start uploading %v from %v", filename, casesType)
	err := uc.repo.UploadCasesFile(problemID, casesFile, filename, casesType)
	if err != nil {
		return err
	}
	uc.log.WithContext(ctx).Infof("Upload %v completed.", filename)
	return nil
}
