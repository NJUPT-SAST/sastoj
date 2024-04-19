package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
)

func (s *UserContestService) GetProblems(ctx context.Context, req *pb.GetProblemsRequest) (*pb.GetProblemsReply, error) {
	rv, err := s.problemUc.ListProblem(ctx, req.ContestId)
	if err != nil {
		return &pb.GetProblemsReply{}, err
	}
	reply := &pb.GetProblemsReply{}
	for _, p := range rv {
		reply.Problems = append(reply.Problems, &pb.GetProblemsReply_Problem{
			Id:    p.ID,
			Title: p.Title,
			Point: int32(p.Point),
			Index: int32(p.Index),
		})
	}
	return reply, nil
}

func (s *UserContestService) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	rv, err := s.problemUc.GetProblem(ctx, req.ProblemId, req.ContestId)
	if err != nil {
		return &pb.GetProblemReply{}, err
	}
	return &pb.GetProblemReply{
		Id:      rv.ID,
		Title:   rv.Title,
		Content: rv.Content,
		Point:   int32(rv.Point),
	}, nil
}

func (s *UserContestService) getProblemCaseVer(ctx context.Context, problemId int64) (int8, error) {
	return s.problemUc.GetProblemCaseVer(ctx, problemId)
}
