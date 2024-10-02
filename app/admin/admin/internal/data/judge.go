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

func (r *judgeRepo) SubmitJudge(ctx context.Context, submissionID int64, point int32) error {
	_, err := r.data.db.Submission.Update().SetState(2).SetPoint(int16(point)).Where(submission.IDEQ(submissionID)).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *judgeRepo) GetJudgableProblems(ctx context.Context, userID int64) ([]*biz.Problem, error) {
	problems, err := r.data.db.Problem.Query().Where(problem.HasAdjudicatorsWith(group.HasUsersWith(user.IDEQ(userID)))).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Problem, 0)
	for _, p := range problems {
		rv = append(rv, &biz.Problem{
			ID:          p.ID,
			Title:       p.Title,
			Content:     p.Content,
			Score:       p.Score,
			ContestID:   p.ContestID,
			CaseVersion: p.CaseVersion,
			Index:       p.Index,
		})
	}
	return rv, nil
}

func (r *judgeRepo) GetSubmissions(ctx context.Context, problemID int64) ([]*biz.Submission, error) {
	submissions, err := r.data.db.Submission.Query().Where(submission.ProblemIDEQ(problemID)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Submission, 0)
	for _, s := range submissions {
		rv = append(rv, &biz.Submission{
			ID:         s.ID,
			Code:       s.Code,
			Status:     int32(s.State),
			Point:      int32(s.Point),
			CreateTime: s.CreateTime,
		})
	}
	return rv, nil
}

func (r *judgeRepo) GetSubmissionsWithStatus(ctx context.Context, problemID int64, status int32) ([]*biz.Submission, error) {
	submissions, err := r.data.db.Submission.Query().Where(submission.And(submission.ProblemIDEQ(problemID), submission.StateEQ(int16(status)))).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Submission, 0)
	for _, s := range submissions {
		rv = append(rv, &biz.Submission{
			ID:         s.ID,
			Code:       s.Code,
			Status:     int32(s.State),
			Point:      int32(s.Point),
			CreateTime: s.CreateTime,
		})
	}
	return rv, nil
}
