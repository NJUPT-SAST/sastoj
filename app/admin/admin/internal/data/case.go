package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mime/multipart"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/pkg/util"
)

type caseRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemCaseRepo .
func NewProblemCaseRepo(data *Data, logger log.Logger) biz.CaseRepo {
	return &caseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *caseRepo) UploadCasesFile(problemId int64, casesFile multipart.File, filename string, casesType string) error {
	err := r.data.fm.DeleteCase(problemId)
	if err != nil {
		return err
	}
	err = r.data.fm.SaveAndExtractCase(casesFile, filename, problemId)
	if err != nil {
		return err
	}
	if casesType == "hydro" {
		err = r.data.fm.HandleHydroCase(problemId)
		if err != nil {
			return err
		}
	}
	config, err := r.data.fm.GetJudgeConfig(problemId)
	if err != nil {
		return err
	}
	err = r.data.fm.CaseCrlfToLf(problemId, config)
	if err != nil {
		return err
	}
	err = util.CalculateScores(config)
	if err != nil {
		return err
	}
	err = r.data.fm.SetJudgeConfig(problemId, config)
	if err != nil {
		return err
	}
	return r.data.fm.CompressAndArchive(problemId)
}
