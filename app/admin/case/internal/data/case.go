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
		ids = append(ids, rc.ID)
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

func (r *caseRepo) UploadCasesFile(problemId int64, casesFile multipart.File, filename string) (util.JudgeConfig, error) {
	base := r.data.problemCasesLocation
	location := base + "/" + strconv.FormatInt(problemId, 10) + "/"
	if _, err := os.Stat(location); err == nil {
		err := os.RemoveAll(location)
		if err != nil {
			return util.JudgeConfig{}, err
		}
	}
	err := os.Mkdir(location, os.ModePerm)
	if err != nil {
		return util.JudgeConfig{}, err
	}
	f, err := os.OpenFile(location+filename, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return util.JudgeConfig{}, err
	}
	defer f.Close()
	_, _ = io.Copy(f, casesFile)

	// decompressed
	err = archiver.Unarchive(location+filename, location)
	if err != nil {
		return util.JudgeConfig{}, err
	}

	// handle files
	dir, err := os.ReadDir(location + "config")
	if err != nil {
		return util.JudgeConfig{}, err
	}
	for _, file := range dir {
		if file.IsDir() {
			//os.Mkdir(location + "testdata", os.ModePerm)
			err := os.Rename(location+"config"+"/"+file.Name()+"/"+"testdata", location+"testdata")
			if err != nil {
				return util.JudgeConfig{}, err
			}
			err = os.RemoveAll(location + "config")
			if err != nil {
				return util.JudgeConfig{}, err
			}
		}
	}

	// unmarshal toml
	tomlText, err := os.ReadFile(location + "testdata" + "/" + "config.toml")
	if err != nil {
		return util.JudgeConfig{}, err
	}
	config, err := util.UnmarshalToml(tomlText)
	if err != nil {
		return util.JudgeConfig{}, err
	}

	// crlf to lf
	caseNum := len(config.Task.Cases)
	type Empty interface{}
	var empty Empty
	sem := make(chan Empty, caseNum)
	for i := 1; i < len(config.Task.Cases)+1; i++ {
		go func(i int) {
			in, err := os.ReadFile(location + "testdata" + "/" + strconv.Itoa(i) + ".in")
			if err != nil {
				sem <- err
				return
			}
			out, err := os.ReadFile(location + "testdata" + "/" + strconv.Itoa(i) + ".ans")
			if err != nil {
				sem <- err
				return
			}
			err = os.WriteFile(location+"testdata"+"/"+strconv.Itoa(i)+".in", []byte(util.Crlf2lf(string(in[:]))), os.ModePerm)
			if err != nil {
				sem <- err
				return
			}
			err = os.WriteFile(location+"testdata"+"/"+strconv.Itoa(i)+".ans", []byte(util.Crlf2lf(string(out[:]))), os.ModePerm)
			if err != nil {
				sem <- err
				return
			}
			sem <- empty
		}(i)
	}
	for i := 0; i < len(config.Task.Cases); i++ {
		select {
		case r := <-sem:
			if err, ok := r.(error); ok {
				return util.JudgeConfig{}, err
			}
		}
	}

	// compressed
	err = archiver.Archive([]string{location + "testdata"}, location+"/"+"testdata.tar.zst")
	if err != nil {
		return util.JudgeConfig{}, err
	}
	return config, nil
}
