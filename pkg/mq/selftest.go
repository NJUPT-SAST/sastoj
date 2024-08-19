package mq

import "fmt"

type SelfTest struct {
	ID         string `json:"id,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	Code       string `json:"code,omitempty"`
	Language   string `json:"language,omitempty"`
	Input      string `json:"input,omitempty"`
	IsCompiled bool   `json:"is_compiled,omitempty"`
	Stdout     string `json:"stdout,omitempty"`
	Stderr     string `json:"stderr,omitempty"`
	Time       uint64 `json:"time,omitempty"`
	Memory     uint64 `json:"memory,omitempty"`
	Token      string `json:"secret,omitempty"`
}

func (s SelfTest) String() string {
	return fmt.Sprintf(
		"SelfTest{ID: %s, UserID: %d, Code: %s, Language: %s, Input: %s, IsCompiled: %t, Stdout: %s, Stderr: %s, Time: %d, Memory: %d, Token: %s}",
		s.ID,
		s.UserID,
		s.Code,
		s.Language,
		s.Input,
		s.IsCompiled,
		s.Stdout,
		s.Stderr,
		s.Time,
		s.Memory,
		s.Token,
	)
}
