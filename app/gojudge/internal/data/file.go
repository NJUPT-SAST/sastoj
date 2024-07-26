package data

import (
	"os"
	u "sastoj/pkg/util"
	"strconv"
)

type FileManage struct {
	FileLocation string
}

func (m *FileManage) GetConfig(problemId int64) (*u.JudgeConfig, error) {
	location := m.FileLocation + "/" + strconv.FormatInt(problemId, 10)
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

func (m *FileManage) FetchCase(problemId int64, fileIn string, fileAns string) (in []byte, ans []byte, err error) {
	location := m.FileLocation + "/" + strconv.FormatInt(problemId, 10) + "/"
	in, err = os.ReadFile(location + "/" + fileIn)
	ans, err = os.ReadFile(location + "/" + fileAns)
	return
}
