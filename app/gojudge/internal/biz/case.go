package biz

import (
	"os"
	d "sastoj/app/gojudge/internal/data"
	u "sastoj/pkg/util"
	"strconv"
)

type CaseManage struct {
	data *d.Data
}

func NewCaseManage(data *d.Data) *CaseManage {
	return &CaseManage{
		data: data,
	}
}

func (m *CaseManage) GetConfig(problemId int64) (*u.JudgeConfig, error) {
	location := m.data.FileLocation + "/" + strconv.FormatInt(problemId, 10)
	toml, err := os.ReadFile(location + "/testdata/config.toml")
	if err != nil {
		return nil, err
	}
	config, err := u.UnmarshalToml(toml)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (m *CaseManage) FetchCase(problemId int64, fileIn string, fileAns string) (in []byte, ans []byte, err error) {
	location := m.data.FileLocation + "/" + strconv.FormatInt(problemId, 10) + "/"
	in, err = os.ReadFile(location + "/" + fileIn)
	ans, err = os.ReadFile(location + "/" + fileAns)
	return
}
