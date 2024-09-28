package util

import "testing"

func TestStringCompareIgnoreLineEndSpaceAndTextEndEnter(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				a: "",
				b: "",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				a: "a",
				b: "a ",
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				a: "a",
				b: "a\n",
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				a: "a",
				b: "a\n ",
			},
			want: true,
		},
		{
			name: "Test 5",
			args: args{
				a: "a",
				b: "a\n\n",
			},
			want: true,
		},
		{
			name: "Test 6",
			args: args{
				a: "a",
				b: "a\n\n ",
			},
			want: true,
		},
		{
			name: "Test 7",
			args: args{
				a: "a",
				b: "a  \n",
			},
			want: true,
		},
		{
			name: "Test 8",
			args: args{
				a: "a",
				b: "a  \n ",
			},
			want: true,
		},
		{
			name: "Test 9",
			args: args{
				a: "a",
				b: "a  \n\n",
			},
			want: true,
		},
		{
			name: "Test 10",
			args: args{
				a: "a\nb\n",
				b: "ab",
			},
			want: false,
		},
		{
			name: "Test 11",
			args: args{
				a: "a\nb\n",
				b: "a\nb",
			},
			want: true,
		},
		{
			name: "Test 12",
			args: args{
				a: "a\nb\n",
				b: "a \nb\n",
			},
			want: true,
		},
		{
			name: "Test 13",
			args: args{
				a: "a\nb\n",
				b: "a \nb",
			},
			want: true,
		},
		{
			name: "Test 14",
			args: args{
				a: "a\nb\n",
				b: "a\nb \n",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringCompareIgnoreLineEndSpaceAndTextEndEnter(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("StringCompareIgnoreLineEndSpaceAndTextEndEnter() = %v, want %v", got, tt.want)
			}
		})
	}
}
