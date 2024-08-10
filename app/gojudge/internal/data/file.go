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
	return u.GetConfig(problemId, m.FileLocation)
}

func (m *FileManage) FetchCase(problemId int64, fileIn string, fileAns string) (in []byte, ans []byte, err error) {
	location := m.FileLocation + "/" + strconv.FormatInt(problemId, 10) + "/testdata/"
	in, err = os.ReadFile(location + "/" + fileIn)
	ans, err = os.ReadFile(location + "/" + fileAns)
	return
}
