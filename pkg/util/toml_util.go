package util

import (
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

func UnmarshalToml(doc []byte) (JudgeConfig, error) {

	var cfg JudgeConfig
	err := toml.Unmarshal(doc, &cfg)
	if err != nil {
		return JudgeConfig{}, err
	}
	return cfg, nil
}
