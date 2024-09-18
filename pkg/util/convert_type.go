package util

import (
	contestApi "sastoj/api/sastoj/admin/contest/service/v1"
	userApi "sastoj/api/sastoj/admin/user/service/v1"
	"sastoj/ent/contest"
	"sastoj/ent/user"
	"time"
)

func ContestStateToInt(state contest.State, start, end time.Time) int32 {
	switch state {
	case contest.StateNORMAL:
		now := time.Now()
		switch {
		case now.Before(start):
			return 0
		case now.After(end):
			return 2
		default:
			return 1
		}
	case contest.StateCANCELLED:
		return 3
	case contest.StateHIDDEN:
		return 4
	case contest.StateDELETED:
		return 5
	default:
		return -1
	}
}

func ContestStateToEnt(state int32) (contest.State, error) {
	switch state {
	case 0, 1, 2:
		return contest.StateNORMAL, nil
	case 3:
		return contest.StateCANCELLED, nil
	case 4:
		return contest.StateHIDDEN, nil
	case 5:
		return contest.StateDELETED, nil
	default:
		return "", contestApi.ErrorContestInvalid("contest state invalid")
	}
}

func UserStateToInt(state user.State) int16 {
	switch state {
	case user.StateNORMAL:
		return 0
	case user.StateBANNED:
		return 1
	case user.StateINACTIVE:
		return 2
	default:
		return -1
	}
}

func UserStateToEnt(state int16) (user.State, error) {
	switch state {
	case 0:
		return user.StateNORMAL, nil
	case 1:
		return user.StateBANNED, nil
	case 2:
		return user.StateINACTIVE, nil
	default:
		return "", userApi.ErrorUserInvalid("user state invalid")
	}
}
