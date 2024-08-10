package util

import (
	"errors"
	"strconv"
)

func CalculateScores(config *JudgeConfig) error {
	switch config.Judge.JudgeType {
	case "classic":
		switch config.Task.TaskType {
		case "simple":
			sum := config.Score
			var num int16 = 0
			for _, c := range config.Task.Cases {
				if c.Score == 0 {
					num += 1
				} else {
					sum -= c.Score
				}
			}
			if sum < 0 {
				return errors.New("the sum of subScore(" + strconv.Itoa(int(config.Score-sum)) + ") is larger than sumScore(" + strconv.Itoa(int(config.Score)) + ")")
			}
			if num > sum {
				return errors.New("the number of subScore(" + strconv.Itoa(int(num)) + ") is larger than sumScore(" + strconv.Itoa(int(config.Score)) + ")")
			}
			avg := sum / num
			remainder := sum % num
			for i := range config.Task.Cases {
				c := &config.Task.Cases[i]
				if c.Score == 0 {
					num--
					c.Score = avg
					if num < remainder {
						c.Score++
					}
				}
			}
			return nil
		case "subtasks":
			sum := config.Score
			var num int16 = 0
			for _, sub := range config.Task.Subtasks {
				if sub.Score == 0 {
					num += 1
				} else {
					sum -= sub.Score
				}
			}
			if sum < 0 {
				return errors.New("the sum of subScore(" + strconv.Itoa(int(config.Score-sum)) + ") is larger than sumScore(" + strconv.Itoa(int(config.Score)) + ")")
			}
			if num > sum {
				return errors.New("the number of subScore(" + strconv.Itoa(int(num)) + ") is larger than sumScore(" + strconv.Itoa(int(config.Score)) + ")")
			}

			avg := sum / num
			remainder := sum % num
			for i := range config.Task.Subtasks {
				sub := &config.Task.Subtasks[i]
				if sub.Score == 0 {
					num--
					sub.Score = avg
					if num < remainder {
						sub.Score++
					}
				}
				caseLen := int16(len(sub.Cases))
				subAvg := sub.Score / caseLen
				subRemainder := sub.Score % caseLen
				for j := range sub.Cases {
					c := &sub.Cases[j]
					if c.Score == 0 {
						caseLen--
						c.Score = subAvg
						if caseLen < subRemainder {
							c.Score++
						}
					}
				}
			}
			return nil
		}
	}
	return errors.New("invalid judge type")
}
