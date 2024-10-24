package service

import (
	"context"
	pb "sastoj/api/sastoj/user/contest/service/v1"
)

func (s *GatewayService) GetProblems(ctx context.Context, req *pb.GetProblemsRequest) (*pb.GetProblemsReply, error) {
	rv, err := s.problemUc.GetProblems(ctx, req.GetContestId())
	if err != nil {
		return nil, err
	}
	reply := &pb.GetProblemsReply{}
	for _, p := range rv {
		reply.Problems = append(reply.Problems, &pb.GetProblemsReply_Problem{
			Id:       p.ID,
			Type:     p.Type,
			Title:    p.Title,
			Score:    p.Score,
			Index:    int32(p.Index),
			Metadata: p.Metadata,
		})
	}
	return reply, nil
}
func (s *GatewayService) GetProblem(ctx context.Context, req *pb.GetProblemRequest) (*pb.GetProblemReply, error) {
	rv, err := s.problemUc.GetProblem(ctx, req.GetContestId(), req.GetProblemId())
	if err != nil {
		return nil, err
	}
	return &pb.GetProblemReply{
		Id:       rv.ID,
		Type:     rv.Type,
		Title:    rv.Title,
		Content:  rv.Content,
		Score:    rv.Score,
		Metadata: rv.Metadata,
	}, nil
}
