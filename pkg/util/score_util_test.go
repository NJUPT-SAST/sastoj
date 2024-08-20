package util

import "testing"

func TestCalculateScores(t *testing.T) {
	type args struct {
		config *JudgeConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantRes []int16
	}{
		{
			name: "Test CalculateScores 1",
			args: args{
				config: &JudgeConfig{
					Judge: Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: Task{
						TaskType: "simple",
						Cases: []Cases{
							{
								Score: 30,
							},
							{
								Score: 40,
							},
							{
								Score: 29,
							},
						},
					},
				},
			},
			wantErr: false,
			wantRes: []int16{30, 40, 30},
		},
		{
			name: "Test CalculateScores 2",
			args: args{
				config: &JudgeConfig{
					Judge: Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: Task{
						TaskType: "simple",
						Cases: []Cases{
							{
								Score: 10,
							},
							{
								Score: 90,
							},
							{
								Score: 10,
							},
						},
					},
				},
			},
			wantErr: true,
			wantRes: []int16{10, 90, 10},
		},
		{
			name: "Test CalculateScores 3",
			args: args{
				config: &JudgeConfig{
					Judge: Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: Task{
						TaskType: "simple",
						Cases: []Cases{
							{
								Score: 20,
							},
							{
								Score: 30,
							},
							{
								Score: 40,
							},
						},
					},
				},
			},
			wantErr: false,
			wantRes: []int16{23, 33, 44},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CalculateScores(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("CalculateScores() error = %v, wantErr %v", err, tt.wantErr)
			}
			for i, c := range tt.args.config.Task.Cases {
				if c.Score != tt.wantRes[i] {
					t.Errorf("CalculateScores() = %v, want %v", c.Score, tt.wantRes[i])
				}
			}
		})
	}
}
