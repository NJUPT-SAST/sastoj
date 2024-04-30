package util

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"strings"
)

type Cases struct {
	Input  string
	Answer string
	Score  int16
}

type Judge struct {
	JudgeType string
	TaskType  string
	Cases     []Cases

	Checker string

	Interactor string
	Subtasks   []Subtasks
}

type Subtasks struct {
	Score int16
	Cases []Cases
}

type ResourceLimits struct {
	Time   int16
	Memory int16
}

type Simple struct {
	Score          int16
	Judge          Judge
	ResourceLimits ResourceLimits
}

func UnmarshalToml(doc []byte) Simple {

	var cfg Simple
	err := toml.Unmarshal(doc, &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("score:", cfg.Score)
	for _, c := range cfg.Judge.Cases {
		fmt.Println("case index: ", strings.Split(c.Input, ".")[0])
	}
	return cfg

	// Output:
	// version: 2
	// name: go-toml
	// tags: [go toml]
}
