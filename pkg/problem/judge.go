package problem

import (
	"errors"
	"sastoj/pkg/util"
)

type Judge interface {
	judging([]bool) ([]int16, error)
}

func Judging(cases []bool, tasks *util.Task) ([]int16, error) {
	var judgeObj Judge
	if len(cases) == 0 {
		return nil, errors.New("no cases found")
	}
	switch tasks.TaskType {
	case "simple":
		judgeObj = &SimpleTask{task: tasks.Cases}
	case "subtasks":
		judgeObj = &SubtaskMin{subtasks: tasks.Subtasks}
	default:
		return nil, errors.New("unknown task type")
	}
	return judgeObj.judging(cases)
}
