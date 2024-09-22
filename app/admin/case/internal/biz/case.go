package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"mime/multipart"
	"sastoj/pkg/util"
	"strconv"
	"strings"
)

type Case struct {
	Id        int64
	ProblemId int64
	Point     int32
	Index     int32
	IsAuto    bool
}

type CaseRepo interface {
	DeleteByProblemId(context.Context, int64) error
	UploadCasesFile(int64, multipart.File, string, string) (*util.JudgeConfig, error)
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
	err := uc.repo.DeleteByProblemId(ctx, pi)
	if err != nil {
		return err
	}
	return nil
}

func (uc *CaseUsecase) UploadCases(ctx context.Context, problemId int64, casesFile multipart.File, filename string, casesType string) error {
	uc.log.WithContext(ctx).Infof("start uploading %v from %v", filename, casesType)
	config, err := uc.repo.UploadCasesFile(problemId, casesFile, filename, casesType)
	if err != nil {
		return err
	}
	uc.log.WithContext(ctx).Infof("Upload %v completed. Start saving info", filename)
	err = uc.repo.DeleteByProblemId(ctx, problemId)
	if err != nil {
		return err
	}
	var cases []*Case
	// TODO: cases problem as special-cases, interactive or subtask type
	switch config.Judge.JudgeType {
	case "classic":
		{
			if config.Task.TaskType == "simple" {
				for _, c := range config.Task.Cases {
					index, err := strconv.Atoi(strings.Split(c.Input, ".")[0])
					if err != nil {
						return err
					}
					cases = append(cases, &Case{
						ProblemId: problemId,
						Point:     int32(c.Score),
						Index:     int32(index),
						IsAuto:    false,
					})
				}
				if err != nil {
					return err
				}
				return nil
			} else if config.Task.TaskType == "subtask" {

			}
			return nil
		}
	case "special-cases":
		return nil
	case "interactive":
		return nil
	}
	err = errors.New("missing cases-type in config.toml")
	return err
}
