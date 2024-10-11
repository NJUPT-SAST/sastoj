package problem

import (
	"sastoj/pkg/file"
	"sastoj/pkg/util"
)

type SubtaskMin struct {
	subtasks file.Subtasks
}

func (s *SubtaskMin) judging(cases []bool) (tasksPoint int16, casesPoint []int16, err error) {
	tasksPoint = s.subtasks.Score
	casesPoint = make([]int16, len(cases), len(cases))
	err = nil
	flag := true
	for _, c := range s.subtasks.Cases {
		i := util.GetCaseIndex(c.Input) - 1
		if !cases[i] {
			casesPoint[i] = 0
			flag = false
		} else {
			casesPoint[i] = c.Score
		}
	}
	if !flag {
		tasksPoint = 0
	}
	return
}
