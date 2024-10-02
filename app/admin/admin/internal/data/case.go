package data

import (
	"mime/multipart"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
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

func (r *caseRepo) UploadCasesFile(problemID int64, casesFile multipart.File, filename string, casesType string) error {
	err := r.data.fm.DeleteCase(problemID)
	if err != nil {
		return err
	}
	err = r.data.fm.SaveAndExtractCase(casesFile, filename, problemID)
	if err != nil {
		return err
	}
	if casesType == "hydro" {
		err = r.data.fm.HandleHydroCase(problemID)
		if err != nil {
			return err
		}
	}
	config, err := r.data.fm.GetJudgeConfig(problemID)
	if err != nil {
		return err
	}
	err = r.data.fm.CaseCrlfToLf(problemID, config)
	if err != nil {
		return err
	}
	err = util.CalculateScores(config)
	if err != nil {
		return err
	}
	err = r.data.fm.SetJudgeConfig(problemID, config)
	if err != nil {
		return err
	}
	return r.data.fm.CompressAndArchive(problemID)
}
