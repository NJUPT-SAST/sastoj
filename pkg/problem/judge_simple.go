package problem

import (
	"errors"
	"sastoj/pkg/util"
)

type SimpleTask struct {
	task []util.Cases
}

func (s *SimpleTask) judging(cases []bool) ([]int16, error) {
	if len(cases) != len(s.task) {
		return nil, errors.New("cases and tasks not match")
	}
	res := make([]int16, len(cases), len(cases))
	for i, v := range cases {
		if v {
			res[i] = s.task[i].Score
		} else {
			res[i] = 0
		}
	}
	return res, nil
}
