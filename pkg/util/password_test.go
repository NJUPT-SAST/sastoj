package util

import "testing"

func TestGenerateRandomString(t *testing.T) {
	type args struct {
		length  int
		charset string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test GenerateRandomString",
			args: args{
				length:  10,
				charset: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateRandomString(tt.args.length, tt.args.charset)
			t.Logf("GenerateRandomString: %v", got)
		})
	}
}
