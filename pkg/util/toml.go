package util

import (
	"os"
	"strconv"

	"github.com/pelletier/go-toml/v2"
)

type Cases struct {
	Input  string
	Answer string
	Score  int16
}

type Judge struct {
	JudgeType  string
	Interactor string
}
type Task struct {
	TaskType string
	Cases    []Cases
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

func ReadTomlFile(problemID int64, baseLocation string) ([]byte, error) {
	tomlFile, err := os.ReadFile(baseLocation + "/" + strconv.FormatInt(problemID, 10) + "/testdata/config.toml")
	if err != nil {
		return nil, err
	}
	return tomlFile, nil
}

func WriteTomlFile(problemID int64, baseLocation string, tomlFile []byte) error {
	err := os.WriteFile(baseLocation+"/"+strconv.FormatInt(problemID, 10)+"/testdata/config.toml", tomlFile, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func GetConfig(problemID int64, baseLocation string) (config *JudgeConfig, err error) {
	tomlFile, err := ReadTomlFile(problemID, baseLocation)
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(tomlFile, &config)
	if err != nil {
		return nil, err
	}
	return
}

func SetConfig(problemID int64, baseLocation string, config *JudgeConfig) error {
	tomlFile, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	return WriteTomlFile(problemID, baseLocation, tomlFile)
}
