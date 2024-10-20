package problem

import (
	"errors"
	"sastoj/pkg/file"
)

type Judge interface {
	judging([]bool) (taskPoint int16, casesPoint []int16, err error)
}

func Judging(cases []bool, taskType string, subtask file.Subtasks) (taskPoint int16, casesPoint []int16, err error) {
	var judgeObj Judge
	if len(cases) == 0 {
		return 0, nil, errors.New("no cases found")
	}
	switch taskType {
	case "simple":
		judgeObj = &SimpleTask{task: subtask}
	case "subtasks":
		judgeObj = &SubtaskMin{subtasks: subtask}
	default:
		return 0, nil, errors.New("unknown task type")
	}
	return judgeObj.judging(cases)
}
