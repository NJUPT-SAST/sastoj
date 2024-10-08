package data

import (
	"context"
	"sastoj/app/admin/admin/internal/biz"
	"sastoj/ent/group"
	"sastoj/ent/problem"
	"sastoj/ent/submission"
	"sastoj/ent/user"

	"github.com/go-kratos/kratos/v2/log"
)

type judgeRepo struct {
	data *Data
	log  *log.Helper
}

// NewJudgeRepo .
func NewJudgeRepo(data *Data, logger log.Logger) biz.JudgeRepo {
	return &judgeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *judgeRepo) SubmitJudge(ctx context.Context, submissionId int64, point int32) error {
	_, err := r.data.db.Submission.Update().SetState(2).SetPoint(int16(point)).Where(submission.IDEQ(submissionId)).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *judgeRepo) GetJudgableProblems(ctx context.Context, userId int64) ([]*biz.Problem, error) {
	problems, err := r.data.db.Problem.Query().Where(problem.HasAdjudicatorsWith(group.HasUsersWith(user.IDEQ(userId)))).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Problem, 0)
	for _, p := range problems {
		rv = append(rv, &biz.Problem{
			Id:          p.ID,
			Title:       p.Title,
			Content:     p.Content,
			Point:       int32(p.Score),
			ContestId:   p.ContestID,
			CaseVersion: int32(p.CaseVersion),
			Index:       int32(p.Index),
			Config:      "",
		})
	}
	return rv, nil
}

func (r *judgeRepo) GetSubmissions(ctx context.Context, problemId int64) ([]*biz.Submission, error) {
	submissions, err := r.data.db.Submission.Query().Where(submission.ProblemIDEQ(problemId)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Submission, 0)
	for _, s := range submissions {
		rv = append(rv, &biz.Submission{
			Id:         s.ID,
			Code:       s.Code,
			Status:     int32(s.State),
			Point:      int32(s.Point),
			CreateTime: s.CreateTime,
		})
	}
	return rv, nil
}

func (r *judgeRepo) GetSubmissionsWithStatus(ctx context.Context, problemId int64, status int32) ([]*biz.Submission, error) {
	submissions, err := r.data.db.Submission.Query().Where(submission.And(submission.ProblemIDEQ(problemId), submission.StateEQ(int16(status)))).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Submission, 0)
	for _, s := range submissions {
		rv = append(rv, &biz.Submission{
			Id:         s.ID,
			Code:       s.Code,
			Status:     int32(s.State),
			Point:      int32(s.Point),
			CreateTime: s.CreateTime,
		})
	}
	return rv, nil
}
