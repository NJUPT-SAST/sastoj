package util

import "testing"

func TestCrlf2lf(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "test1",
			args: "a\r\nb\nc\r\n",
			want: "a\nb\nc\n",
		},
		{
			name: "test2",
			args: "a\r\nb\nc\r",
			want: "a\nb\nc\r",
		},
		{
			name: "test3",
			args: "a3\\r\nb\nc\r\n",
			want: "a3\\r\nb\nc\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Crlf2lf(tt.args); got != tt.want {
				t.Errorf("Crlf2lf() = %v, want %v", got, tt.want)
			}
		})
	}
}
