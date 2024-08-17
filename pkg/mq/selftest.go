package mq

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
