package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"os"
	"sastoj/app/rsjudge/cases/internal/biz"
	"strconv"
)

type casesRepo struct {
	data *Data
	log  *log.Helper
}

// NewJudgeRepo .
func NewJudgeRepo(data *Data, logger log.Logger) biz.CasesRepo {
	return &casesRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r casesRepo) FetchCases(problemId int64) ([]byte, int64, error) {
	base := r.data.problemCasesLocation
	location := base + "/" + strconv.FormatInt(problemId, 10) + "/"
	tomlText, err := os.ReadFile(location + "testdata.tar.zst")
	if err != nil {
		return nil, 0, err
	}
	chunkSize := r.data.chunkSize
	return tomlText, chunkSize, nil
}
