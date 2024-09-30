package mq

import (
	"fmt"
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
	TotalTime  uint64    `json:"total_time,omitempty"`
	MaxMemory  uint64    `json:"max_memory,omitempty"`
	Language   string    `json:"language,omitempty"`
	Stderr     string    `json:"stderr,omitempty"`
	CaseVer    int16     `json:"case_version,omitempty"`
	Token      string    `json:"secret,omitempty"`
}

func (s Submission) String() string {
	return fmt.Sprintf(
		"Submission{ID: %s, UserId: %d, ProblemID: %d, Code: %s, Status: %d, Point: %d, CreateTime: %s, TotalTime: %d, MaxMemory: %d, Language: %s, Stderr: %s, Token: %s}",
		s.ID,
		s.UserID,
		s.ProblemID,
		s.Code,
		s.Status,
		s.Point,
		s.CreateTime,
		s.TotalTime,
		s.MaxMemory,
		s.Language,
		s.Stderr,
		s.Token,
	)
}
