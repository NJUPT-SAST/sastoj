package mq

import (
	"sastoj/ent"
	"strconv"
	"time"
)

type Submission struct {
	ID         string    `json:"id,omitempty"`
	UserID     int64     `json:"user_id,omitempty"`
	ProblemID  int64     `json:"problem_id,omitempty"`
	Code       string    `json:"code,omitempty"`
	Status     int16     `json:"state,omitempty"`
	Point      int16     `json:"point,omitempty"`
	CreateTime time.Time `json:"create_time"`
	TotalTime  int32     `json:"total_time,omitempty"`
	MaxMemory  int32     `json:"max_memory,omitempty"`
	Language   string    `json:"language,omitempty"`
	Token      string    `json:"secret,omitempty"`
}

func Ent2mq(e *ent.Submission) *Submission {
	return &Submission{
		ID:         strconv.FormatInt(e.ID, 10),
		UserID:     e.UserID,
		ProblemID:  e.ProblemID,
		Code:       e.Code,
		Status:     e.Status,
		Point:      e.Point,
		CreateTime: e.CreateTime,
		TotalTime:  e.TotalTime,
		MaxMemory:  e.MaxMemory,
		Language:   e.Language,
		Token:      "",
	}
}
