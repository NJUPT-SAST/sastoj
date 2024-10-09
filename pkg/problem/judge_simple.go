package problem

import (
	"errors"
	"sastoj/pkg/file"
	"sastoj/pkg/util"
)

type SimpleTask struct {
	task file.Subtasks
}

func (s *SimpleTask) judging(cases []bool) (taskPoint int16, casesPoint []int16, err error) {
	if len(cases) != len(s.task.Cases) {
		return 0, nil, errors.New("cases and tasks not match")
	}
	res := make([]int16, len(cases), len(cases))
	taskPoint = 0
	for _, c := range s.task.Cases {
		i := util.GetCaseIndex(c.Input) - 1
		if cases[i] {
			res[i] = c.Score
			taskPoint += c.Score
		} else {
			res[i] = 0
		}
	}
	return taskPoint, res, nil
}
