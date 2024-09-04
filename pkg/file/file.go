package file

import (
	"os"
	u "sastoj/pkg/util"
	"strconv"
)

type Manager struct {
	location string
}

// GetJudgeConfig get judge config from file system
func (m *Manager) GetJudgeConfig(problemId int64) (*u.JudgeConfig, error) {
	return u.GetConfig(problemId, m.location)
}

// FetchCase fetch test case from file system
func (m *Manager) FetchCase(problemId int64, fileIn string, fileAns string) (in []byte, ans []byte, err error) {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/testdata/"
	in, err = os.ReadFile(location + "/" + fileIn)
	ans, err = os.ReadFile(location + "/" + fileAns)
	return
}

// NewManager create a new file manager
func NewManager(fileLocation string) *Manager {
	return &Manager{location: fileLocation}
}
