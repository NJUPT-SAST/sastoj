package file

import (
	"io"
	"mime/multipart"
	"os"
	u "sastoj/pkg/util"
	"strconv"
	"sync"
)

type Manager struct {
	location string
}

// GetJudgeConfig get judge config from file system
func (m *Manager) GetJudgeConfig(problemId int64) (*u.JudgeConfig, error) {
	return u.GetConfig(problemId, m.location)
}

// SetJudgeConfig save judge config to file system
func (m *Manager) SetJudgeConfig(problemId int64, config *u.JudgeConfig) error {
	return u.SetConfig(problemId, m.location, config)
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

// DeleteCase delete case from file system
func (m *Manager) DeleteCase(problemId int64) error {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/"
	if _, err := os.Stat(location); err == nil {
		err := os.RemoveAll(location)
		if err != nil {
			return err
		}
	}
	return nil
}

// HandleHydroCase handle hydro case from file system
func (m *Manager) HandleHydroCase(problemId int64) error {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/"
	dir, err := os.ReadDir(location + "config")
	if err != nil {
		return err
	}
	for _, file := range dir {
		if file.IsDir() {
			err := os.Rename(location+"config"+"/"+file.Name()+"/"+"testdata", location+"testdata")
			if err != nil {
				return err
			}
			err = os.RemoveAll(location + "config")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CompressAndArchive compress and archive case from file system
func (m *Manager) CompressAndArchive(problemId int64) error {
	return u.CompressAndArchive(m.location+"/"+strconv.FormatInt(problemId, 10)+"testdata"+"/", ".tar.zst")
}

// CaseCrlfToLf set .in and .out files crlf to lf
func (m *Manager) CaseCrlfToLf(problemId int64, config *u.JudgeConfig) error {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/"
	wg := sync.WaitGroup{}
	wg.Add(len(config.Task.Cases))
	errChannel := make(chan error)
	for _, c := range config.Task.Cases {
		go func() {
			defer wg.Done()
			in, err := os.ReadFile(location + "testdata" + "/" + c.Input)
			if err != nil {
				errChannel <- err
				return
			}
			out, err := os.ReadFile(location + "testdata" + "/" + c.Answer)
			if err != nil {
				errChannel <- err
				return
			}
			err = os.WriteFile(location+"testdata"+"/"+c.Input, []byte(u.Crlf2lf(string(in[:]))), os.ModePerm)
			if err != nil {
				errChannel <- err
				return
			}
			err = os.WriteFile(location+"testdata"+"/"+c.Answer, []byte(u.Crlf2lf(string(out[:]))), os.ModePerm)
			if err != nil {
				errChannel <- err
				return
			}
		}()
	}
	wg.Wait()
	if len(errChannel) != 0 {
		return <-errChannel
	}
	return nil
}

// SaveAndExtractCase save and extract case file to file system
func (m *Manager) SaveAndExtractCase(casesFile multipart.File, filename string, problemId int64) error {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/"
	err := os.Mkdir(location, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(location+filename, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, casesFile)
	if err != nil {
		return err
	}
	err = u.ExtractExc(f, location, []string{"config.yaml"})
	if err != nil {
		return err
	}
	return nil
}
