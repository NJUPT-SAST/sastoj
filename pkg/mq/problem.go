package mq

type Problem struct {
	ID       int64             `json:"id,omitempty"`
	Title    string            `json:"title,omitempty"`
	Type     string            `json:"type_id,omitempty"`
	Content  string            `json:"content,omitempty"`
	Score    int16             `json:"score,omitempty"`
	Index    int16             `json:"index,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}
