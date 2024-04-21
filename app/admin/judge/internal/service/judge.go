package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sastoj/app/admin/judge/internal/biz"

	pb "sastoj/api/sastoj/admin/judge/service/v1"
)

type JudgeService struct {
	pb.UnimplementedJudgeServer
	uc *biz.JudgeUsecase
}

func NewJudgeService(usecase *biz.JudgeUsecase) *JudgeService {
	return &JudgeService{
		uc: usecase,
	}
}

func (s *JudgeService) SubmitJudge(ctx context.Context, req *pb.SubmitJudgeRequest) (*pb.SubmitJudgeReply, error) {
	err := s.uc.SubmitJudge(ctx, req.SubmitId, req.Point)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (s *JudgeService) GetJudgableProblems(ctx context.Context, req *pb.GetJudgableProblemsRequest) (*pb.GetJudgableProblemsReply, error) {
	ps, err := s.uc.GetJudgableProblems(ctx, 1)
	if err != nil {
		return nil, err
	}
	reply := &pb.GetJudgableProblemsReply{}
	for _, p := range ps {
		reply.Results = append(reply.Results, &pb.Problem{
			Id:          p.Id,
			Title:       p.Title,
			Content:     p.Content,
			Point:       p.Point,
			ContestId:   p.ContestId,
			CaseVersion: p.CaseVersion,
			Index:       p.Index,
			Config:      p.Config,
		})
	}
	return reply, nil
}
func (s *JudgeService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionsRequest) (*pb.GetSubmissionsReply, error) {
	ps, err := s.uc.GetSubmissions(ctx, req.GetProblemId(), req.GetStatus())
	if err != nil {
		return nil, err
	}
	reply := &pb.GetSubmissionsReply{}
	for _, p := range ps {
		reply.Submissions = append(reply.Submissions, &pb.Submission{
			Id:         p.Id,
			Code:       p.Code,
			Point:      p.Point,
			CreateTime: timestamppb.New(p.CreateTime),
		})
	}
	return reply, nil
}
