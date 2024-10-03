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

func (r *caseRepo) UploadCasesFile(problemId int64, casesFile multipart.File, filename string, casesType string) error {
	err := r.data.jcm.DeleteCase(problemId)
	if err != nil {
		return err
	}
	err = r.data.jcm.SaveAndExtractCase(casesFile, filename, problemId)
	if err != nil {
		return err
	}
	if casesType == "hydro" {
		err = r.data.jcm.HandleHydroCase(problemId)
		if err != nil {
			return err
		}
	}
	config, err := r.data.jcm.GetConfig(problemId)
	if err != nil {
		return err
	}
	err = r.data.jcm.CaseCrlfToLf(problemId, config)
	if err != nil {
		return err
	}
	err = util.CalculateScores(config)
	if err != nil {
		return err
	}
	err = r.data.jcm.SetConfig(problemId, config)
	if err != nil {
		return err
	}
	return r.data.jcm.CompressAndArchive(problemId)
}
