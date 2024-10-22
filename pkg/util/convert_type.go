package util

import (
	api "sastoj/api/sastoj/admin/admin/service/v1"
	pb "sastoj/api/sastoj/admin/admin/service/v1"
	"sastoj/ent/contest"
	"sastoj/ent/problem"
	"sastoj/ent/user"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
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
		return "", api.ErrorContestInvalid("contest state invalid")
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
		return "", api.ErrorUserInvalid("user state invalid")
	}
}

func VisToPb(v problem.Visibility) pb.Visibility {
	switch v {
	case problem.VisibilityPRIVATE:
		return pb.Visibility_Private
	case problem.VisibilityPUBLIC:
		return pb.Visibility_Public
	case problem.VisibilityCONTEST:
		return pb.Visibility_Contest
	default:
		return pb.Visibility_Private
	}
}

func VisToEnt(v pb.Visibility) problem.Visibility {
	switch v {
	case pb.Visibility_Private:
		return problem.VisibilityPRIVATE
	case pb.Visibility_Public:
		return problem.VisibilityPUBLIC
	case pb.Visibility_Contest:
		return problem.VisibilityCONTEST
	default:
		return problem.VisibilityPRIVATE
	}
}

func ConvertTimeToTimeStamp(tm time.Time) *timestamppb.Timestamp {

	return &timestamppb.Timestamp{
		Seconds: tm.Unix(),
		Nanos:   int32(tm.Nanosecond()),
	}
}
