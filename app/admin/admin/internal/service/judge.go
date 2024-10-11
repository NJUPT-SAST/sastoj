package service

import (
	"context"
	pb "sastoj/api/sastoj/admin/admin/service/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *AdminService) SubmitJudge(ctx context.Context, req *pb.SubmitJudgeRequest) (*pb.SubmitJudgeReply, error) {
	err := s.jc.SubmitJudge(ctx, req.SubmissionId, req.Point)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
func (s *AdminService) GetJudgableProblems(ctx context.Context, req *pb.GetJudgableProblemsRequest) (*pb.GetJudgableProblemsReply, error) {
	ps, err := s.jc.GetJudgableProblems(ctx, 1)
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
func (s *AdminService) GetSubmissions(ctx context.Context, req *pb.GetSubmissionsRequest) (*pb.GetSubmissionsReply, error) {
	ps, err := s.jc.GetSubmissions(ctx, req.GetProblemId(), req.GetStatus())
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

func (s *AdminService) GetReferenceAnswer(ctx context.Context, req *pb.GetReferenceAnswerRequest) (*pb.GetReferenceAnswerReply, error) {
	answer, err := s.jc.GetReferenceAnswer(ctx, req.GetProblemId())
	if err != nil {
		return nil, err
	}
	reply := &pb.GetReferenceAnswerReply{
		Answer: *answer,
	}
	return reply, nil
}
