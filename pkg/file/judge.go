package file

import (
	"github.com/pelletier/go-toml/v2"
	"io"
	"mime/multipart"
	"os"
	u "sastoj/pkg/util"
	"strconv"
	"sync"
)

type JudgeConfigManager struct {
	BaseConfigManager
}

type Cases struct {
	Input  string
	Answer string
	Score  int16
}

type Judge struct {
	JudgeType string
}

type Task struct {
	TaskType string
	Checker  string
	Subtasks []Subtasks
}

type Subtasks struct {
	Score int16
	Cases []Cases
}

type ResourceLimits struct {
	Time   int16
	Memory int32
}

type JudgeConfig struct {
	Score          int16
	Judge          Judge
	ResourceLimits ResourceLimits
	Task           Task
}

// FetchCase fetch test case from file system
func (m *JudgeConfigManager) FetchCase(problemId int64, fileIn string, fileAns string) (in []byte, ans []byte, err error) {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/testdata/"
	in, err = os.ReadFile(location + "/" + fileIn)
	ans, err = os.ReadFile(location + "/" + fileAns)
	return
}

// DeleteCase delete case from file system
func (m *JudgeConfigManager) DeleteCase(problemId int64) error {
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
func (m *JudgeConfigManager) HandleHydroCase(problemId int64) error {
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
func (m *JudgeConfigManager) CompressAndArchive(problemId int64) error {
	return u.CompressAndArchive(m.location+"/"+strconv.FormatInt(problemId, 10)+"testdata"+"/", ".tar.zst")
}

// CaseCrlfToLf set .in and .out files crlf to lf
func (m *JudgeConfigManager) CaseCrlfToLf(problemId int64, config *JudgeConfig) error {
	location := m.location + "/" + strconv.FormatInt(problemId, 10) + "/"
	wg := sync.WaitGroup{}
	for _, subtask := range config.Task.Subtasks {
		wg.Add(len(subtask.Cases))
		errChannel := make(chan error)
		for _, c := range subtask.Cases {
			c := c
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
	}
	return nil
}

// SaveAndExtractCase save and extract case file to file system
func (m *JudgeConfigManager) SaveAndExtractCase(casesFile multipart.File, filename string, problemId int64) error {
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

// NewJudgeConfigManager create a new file manager
func NewJudgeConfigManager(fileLocation string) *JudgeConfigManager {
	return &JudgeConfigManager{
		BaseConfigManager: BaseConfigManager{location: fileLocation},
	}
}

func (m *JudgeConfigManager) GetConfig(problemId int64) (*JudgeConfig, error) {
	tomlFile, err := m.ReadFile(problemId)
	if err != nil {
		return nil, err
	}

	var config JudgeConfig
	err = toml.Unmarshal(tomlFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (m *JudgeConfigManager) SetConfig(problemId int64, config *JudgeConfig) error {
	tomlFile, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return m.WriteFile(problemId, tomlFile)
}
