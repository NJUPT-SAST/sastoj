package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"mime/multipart"
	"os"
	"sastoj/app/admin/case/internal/biz"
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

func (r *caseRepo) DeleteByProblemId(ctx context.Context, pi int64) error {
	return nil
}

func (r *caseRepo) UploadCasesFile(problemId int64, casesFile multipart.File, filename string, casesType string) (*util.JudgeConfig, error) {
	baseLocation := r.data.problemCasesLocation
	location := baseLocation + "/" + strconv.FormatInt(problemId, 10) + "/"
	if _, err := os.Stat(location); err == nil {
		err := os.RemoveAll(location)
		if err != nil {
			return nil, err
		}
	}
	err := os.Mkdir(location, os.ModePerm)
	if err != nil {
		return nil, err
	}
	f, err := os.OpenFile(location+filename, os.O_RDWR|os.O_CREATE, 0o666)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	_, err = io.Copy(f, casesFile)
	if err != nil {
		return nil, err
	}

	// handle files
	err = util.ExtractExc(f, location, []string{"config.yaml"})
	if err != nil {
		return nil, err
	}
	if casesType == "hydro" {
		dir, err := os.ReadDir(location + "config")
		if err != nil {
			return nil, err
		}
		for _, file := range dir {
			if file.IsDir() {
				//os.Mkdir(location + "testdata", os.ModePerm)
				err := os.Rename(location+"config"+"/"+file.Name()+"/"+"testdata", location+"testdata")
				if err != nil {
					return nil, err
				}
				err = os.RemoveAll(location + "config")
				if err != nil {
					return nil, err
				}
			}
		}
	}

	// unmarshal toml
	config, err := util.GetConfig(problemId, baseLocation)
	if err != nil {
		return nil, err
	}

	// crlf to lf
	caseNum := len(config.Task.Cases)
	type Empty interface{}
	var empty Empty
	sem := make(chan Empty, caseNum)
	for i := 0; i < len(config.Task.Cases); i++ {
		go func(i int) {
			in, err := os.ReadFile(location + "testdata" + "/" + config.Task.Cases[i].Input)
			if err != nil {
				sem <- err
				return
			}
			out, err := os.ReadFile(location + "testdata" + "/" + config.Task.Cases[i].Answer)
			if err != nil {
				sem <- err
				return
			}
			err = os.WriteFile(location+"testdata"+"/"+config.Task.Cases[i].Input, []byte(util.Crlf2lf(string(in[:]))), os.ModePerm)
			if err != nil {
				sem <- err
				return
			}
			err = os.WriteFile(location+"testdata"+"/"+config.Task.Cases[i].Answer, []byte(util.Crlf2lf(string(out[:]))), os.ModePerm)
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
				return nil, err
			}
		}
	}

	// count scores
	err = util.CalculateScores(config)
	if err != nil {
		return nil, err
	}

	// save toml config to file
	err = util.SetConfig(problemId, baseLocation, *config)
	if err != nil {
		return nil, err
	}

	// compressed
	err = util.CompressAndArchive(location+"testdata", ".tar.zst")
	if err != nil {
		return nil, err
	}
	return config, nil
}
