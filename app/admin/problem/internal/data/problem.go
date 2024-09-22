package data

import (
	"context"
	pb "sastoj/api/sastoj/admin/problem/service/v1"
	"sastoj/ent/problem"
	"sastoj/pkg/middleware/auth"

	"github.com/go-kratos/kratos/v2/log"
)

type ProblemRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemRepo problem
func NewProblemRepo(data *Data, logger log.Logger) *ProblemRepo {
	return &ProblemRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *ProblemRepo) Save(ctx context.Context, g *pb.CreateProblemRequest) (*pb.CreateProblemReply, error) {
	res, err := r.data.db.Problem.Create().
		SetProblemTypeID(g.TypeId).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetScore(int16(g.Point)).
		SetContestID(g.ContestId).
		SetCaseVersion(1).
		SetIndex(int16(g.Index)).
		SetOwnerID(getUserID(ctx)).
		SetVisibility(pb2vis(g.Visibility)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &pb.CreateProblemReply{Id: res.ID}, nil
}

func (r *ProblemRepo) Update(ctx context.Context, g *pb.UpdateProblemRequest) (*int, error) {
	res, err := r.data.db.Problem.Update().
		SetProblemTypeID(g.TypeId).
		SetTitle(g.Title).
		SetContent(g.Content).
		SetScore(int16(g.Point)).
		SetContestID(g.ContestId).
		SetCaseVersion(int16(g.CaseVersion)).
		SetIndex(int16(g.Index)).
		SetOwnerID(getUserID(ctx)).
		SetVisibility(pb2vis(g.Visibility)).
		Where(problem.ID(g.Id)).
		Where(problem.IsDeleted(false)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) FindByID(ctx context.Context, id int64) (*pb.GetProblemReply, error) {
	p, err := r.data.db.Problem.Query().
		Where(problem.ID(id)).
		Where(problem.IsDeleted(false)).
		WithOwner().
		First(ctx)
	if err != nil {
		return nil, err
	}
	owner, err := p.QueryOwner().First(ctx)
	if err != nil {
		return nil, err
	}
	vis := vis2pb(p.Visibility)
	return &pb.GetProblemReply{
		Id:          p.ID,
		Title:       p.Title,
		Content:     p.Content,
		Point:       int32(p.Score),
		ContestId:   p.ContestID,
		CaseVersion: int32(p.CaseVersion),
		Index:       int32(p.Index),
		OwnerId:     owner.ID,
		Visibility:  vis,
	}, nil
}

func (r *ProblemRepo) Delete(ctx context.Context, id int64) (*int, error) {
	res, err := r.data.db.Problem.Update().
		SetIsDeleted(true).
		Where(problem.ID(id)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ProblemRepo) ListPages(ctx context.Context, currency int32, size int32) ([]*pb.ListProblemReply_Problem, error) {
	res, err := r.data.db.Problem.Query().
		Limit(int(size)).Offset(int((currency - 1) * size)).
		WithOwner().
		All(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.ListProblemReply_Problem, 0)
	for _, v := range res {
		vis := vis2pb(v.Visibility)
		owner, err := v.QueryOwner().First(ctx)
		if err != nil {
			return nil, err
		}
		list = append(list, &pb.ListProblemReply_Problem{
			Id:          v.ID,
			TypeId:      v.ProblemTypeID,
			Title:       v.Title,
			Content:     v.Content,
			Point:       int32(v.Score),
			ContestId:   v.ContestID,
			CaseVersion: int32(v.CaseVersion),
			Index:       int32(v.Index),
			OwnerId:     owner.ID,
			Visibility:  vis,
		})
	}
	return list, nil
}

func vis2pb(v problem.Visibility) pb.Visibility {
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

func pb2vis(v pb.Visibility) problem.Visibility {
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

func getUserID(ctx context.Context) int64 {
	claim := ctx.Value("userInfo").(*auth.Claims)
	return claim.UserId
}
