package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mholt/archiver/v3"
	"io"
	"mime/multipart"
	"os"
	"sastoj/app/admin/case/internal/biz"
	"sastoj/ent"
	"sastoj/ent/problemcase"
	"sastoj/pkg/util"
	"strconv"
)

type caseRepo struct {
	data *Data
	log  *log.Helper
}

// NewProblemCaseRepo .
func NewProblemCaseRepo(data *Data, logger log.Logger) biz.CaseRepo {
	return &caseRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *caseRepo) Save(ctx context.Context, pi int64, cs []*biz.Case) ([]int64, error) {
	rcs, err := r.data.db.ProblemCase.MapCreateBulk(cs, func(c *ent.ProblemCaseCreate, i int) {
		c.SetPoint(int16(cs[i].Point)).SetIndex(int16(cs[i].Index)).SetIsAuto(cs[i].IsAuto).SetProblemID(pi)
	}).Save(ctx)
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, rc := range rcs {
		ids = append(ids, int64(rc.ID))
	}
	return ids, nil
}

func (r *caseRepo) Update(ctx context.Context, c *biz.Case) error {
	_, err := r.data.db.ProblemCase.UpdateOneID(c.Id).SetPoint(int16(c.Point)).SetIndex(int16(c.Index)).SetIsAuto(c.IsAuto).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) DeleteByCaseIds(ctx context.Context, cis []int64) error {
	intSlice := make([]int64, len(cis))
	for i, ci := range cis {
		intSlice[i] = ci
	}
	_, err := r.data.db.ProblemCase.Update().Where(
		problemcase.IDIn(intSlice...)).SetIsDeleted(true).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) DeleteByProblemId(ctx context.Context, pi int64) error {
	_, err := r.data.db.ProblemCase.Update().Where(
		problemcase.ProblemIDEQ(pi)).SetIsDeleted(true).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *caseRepo) FindByProblemId(ctx context.Context, pi int64) ([]*biz.Case, error) {
	problemCases, err := r.data.db.ProblemCase.Query().Where(problemcase.ProblemIDEQ(pi)).All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Case, 0)
	for _, p := range problemCases {
		rv = append(rv, &biz.Case{
			Id:     p.ID,
			Point:  int32(p.Point),
			Index:  int32(p.Index),
			IsAuto: p.IsAuto,
		})
	}
	return rv, nil
}

func (r *caseRepo) UploadCasesFile(problemId int64, casesFile multipart.File, filename string) (util.Simple, error) {
	base := r.data.uploadCasesLocation
	location := base + "/" + strconv.FormatInt(problemId, 10) + "/"
	if _, err := os.Stat(location); err == nil {
		err := os.RemoveAll(location)
		if err != nil {
			return util.Simple{}, err
		}
	}
	err := os.Mkdir(location, os.ModePerm)
	if err != nil {
		return util.Simple{}, err
	}
	f, err := os.OpenFile(location+filename, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return util.Simple{}, err
	}
	defer f.Close()
	_, _ = io.Copy(f, casesFile)

	// decompressed
	err = archiver.Unarchive(location+filename, location)
	if err != nil {
		return util.Simple{}, err
	}

	// unmarshal toml
	tomlText, err := os.ReadFile(location + "testdata" + "/" + "config.toml")
	if err != nil {
		return util.Simple{}, err
	}
	config := util.UnmarshalToml(tomlText)
	return config, nil
}
