package util

import (
	"errors"
)

func CalculateScores(config *JudgeConfig) error {
	if config.Score <= 0 {
		return errors.New("problem score should be positive")
	}
	switch config.Judge.JudgeType {
	case "classic":
		switch config.Task.TaskType {
		case "simple":
			count := int16(len(config.Task.Cases))
			if count > config.Score {
				return errors.New("the number of cases is larger than score")
			}
			if count == 0 {
				return errors.New("the number of cases is zero")
			}
			restScore := config.Score
			zeroScoreSum := int16(0)
			for _, c := range config.Task.Cases {
				restScore -= c.Score
				if c.Score == 0 {
					zeroScoreSum++
				}
			}
			if restScore == 0 && zeroScoreSum == 0 {
				return nil
			} else if (restScore == 0 && zeroScoreSum > 0) || restScore < 0 ||
				(restScore > 0 && zeroScoreSum > 0 && restScore < zeroScoreSum) {
				for i := range config.Task.Cases {
					config.Task.Cases[i].Score = 0
				}
				restScore = config.Score
				zeroScoreSum = int16(len(config.Task.Cases))
			} else if restScore > 0 && zeroScoreSum == 0 {
				base := restScore / count
				start := int(count - restScore%count)
				for i, c := range config.Task.Cases {
					score := c.Score + base
					if i >= start {
						score++
					}
					config.Task.Cases[i].Score = score
				}
				return nil
			}
			for i, c := range config.Task.Cases {
				if c.Score == 0 {
					config.Task.Cases[i].Score = restScore / zeroScoreSum
					restScore -= config.Task.Cases[i].Score
					zeroScoreSum--
				}
			}
			return nil
		case "subtasks":
			if len(config.Task.Subtasks) == 0 {
				return errors.New("the number of subtasks is zero")
			}
			if config.Score < int16(len(config.Task.Subtasks)) {
				return errors.New("the number of subtasks is larger than score")
			}
			sum := config.Score
			var num int16 = 0
			for _, sub := range config.Task.Subtasks {
				if sub.Score == 0 {
					num += 1
				}
				sum -= sub.Score
			}
			if sum == 0 && num == 0 {
				return nil
			} else if (sum == 0 && num > 0) || sum < 0 || (sum > 0 && num > 0 && sum < num) {
				for i := range config.Task.Subtasks {
					config.Task.Subtasks[i].Score = 0
				}
				sum = config.Score
				num = int16(len(config.Task.Subtasks))
			} else if sum > 0 && num == 0 {
				num = int16(len(config.Task.Subtasks))
				base := sum / num
				start := int(num - sum%num)
				for i, sub := range config.Task.Subtasks {
					score := sub.Score + base
					if i >= start {
						score++
					}
					config.Task.Subtasks[i].Score = score
				}
				return nil
			}
			for i, sub := range config.Task.Subtasks {
				if sub.Score == 0 {
					config.Task.Subtasks[i].Score = sum / num
					sum -= config.Task.Subtasks[i].Score
					num--
				}
			}
			return nil
		}
	}
	return errors.New("invalid judge type")
}
