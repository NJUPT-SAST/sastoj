package biz

import (
	"sastoj/app/gojudge/internal/conf"
	data "sastoj/app/gojudge/internal/data"
)

type Judge struct {
	data    *data.Data
	Gojudge []Gojudge
}

func NewJudge(data *data.Data) *Judge {
	res := &Judge{data: data, Gojudge: []Gojudge{}}
	return res
}

func (j *Judge) init() {
	for _, middleware := range j.data.JudgeMiddleware {
		switch middleware.Type {
		case conf.JudgeType_gojudge:
			j.Gojudge = append(j.Gojudge, Gojudge{
				Endpoint: middleware.Endpoint,
				Db:       j.data,
			})
		case conf.JudgeType_rsjudge:
		}
	}
}
