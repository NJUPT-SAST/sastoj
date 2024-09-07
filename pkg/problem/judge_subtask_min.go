package problem

import (
	"sastoj/pkg/util"
)

type SubtaskMin struct {
	subtasks []util.Subtasks
}

func (s *SubtaskMin) judging(cases []bool) ([]int16, error) {
	res := make([]int16, len(s.subtasks), len(s.subtasks))

	for i, subtasks := range s.subtasks {
		score := subtasks.Score
		for _, c := range subtasks.Cases {
			index := util.GetCaseIndex(c.Input)
			if !cases[index-1] {
				score = 0
				break
			}
		}
		res[i] = score
	}

	return res, nil
}
