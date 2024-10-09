package util

import (
	"sastoj/pkg/file"
	"testing"
)

func TestCalculateSimpleScores(t *testing.T) {
	type args struct {
		config *file.JudgeConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantRes []int16
	}{
		{
			name: "Test CalculateSimpleScores 1",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
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
			name: "Test CalculateSimpleScores 2",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
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
			wantErr: false,
			wantRes: []int16{33, 33, 34},
		},
		{
			name: "Test CalculateSimpleScores 3",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
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
		{
			name: "Test CalculateSimpleScores 4",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
							{
								Score: 0,
							},
							{
								Score: 0,
							},
							{
								Score: 40,
							},
						},
					},
				},
			},
			wantErr: false,
			wantRes: []int16{30, 30, 40},
		},
		{
			name: "Test CalculateSimpleScores 5",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "simple",
						Cases:    []file.Cases{},
					},
				},
			},
			wantErr: true,
			wantRes: []int16{},
		},
		{
			name: "Test CalculateSimpleScores 6",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 0,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
							{
								Score: 20,
							},
						},
					},
				},
			},
			wantErr: true,
			wantRes: []int16{20},
		},
		{
			name: "Test CalculateSimpleScores 7",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 2,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
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
			wantErr: true,
			wantRes: []int16{20, 30, 40},
		},
		{
			name: "Test CalculateSimpleScores 8",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "simple",
						Cases: []file.Cases{
							{
								Score: 0,
							},
							{
								Score: 0,
							},
							{
								Score: 99,
							},
						},
					},
				},
			},
			wantErr: false,
			wantRes: []int16{33, 33, 34},
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

func TestCalculateSubtaskScores(t *testing.T) {
	type args struct {
		config *file.JudgeConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantRes []int16
	}{
		{
			name: "Test CalculateSubtaskScores 1",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
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
			name: "Test CalculateSubtaskScores 2",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
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
			wantErr: false,
			wantRes: []int16{33, 33, 34},
		},
		{
			name: "Test CalculateSubtaskScores 3",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
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
		{
			name: "Test CalculateSubtaskScores 4",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
							{
								Score: 0,
							},
							{
								Score: 0,
							},
							{
								Score: 40,
							},
						},
					},
				},
			},
			wantErr: false,
			wantRes: []int16{30, 30, 40},
		},
		{
			name: "Test CalculateSubtaskScores 5",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{},
					},
				},
			},
			wantErr: true,
			wantRes: []int16{},
		},
		{
			name: "Test CalculateSubtaskScores 6",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 0,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
							{
								Score: 20,
							},
						},
					},
				},
			},
			wantErr: true,
			wantRes: []int16{20},
		},
		{
			name: "Test CalculateSubtaskScores 7",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 2,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
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
			wantErr: true,
			wantRes: []int16{20, 30, 40},
		},
		{
			name: "Test CalculateSubtaskScores 8",
			args: args{
				config: &file.JudgeConfig{
					Judge: file.Judge{
						JudgeType: "classic",
					},
					Score: 100,
					Task: file.Task{
						TaskType: "subtasks",
						Subtasks: []file.Subtasks{
							{
								Score: 0,
							},
							{
								Score: 0,
							},
							{
								Score: 99,
							},
						},
					},
				},
			},
			wantErr: false,
			wantRes: []int16{33, 33, 34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CalculateScores(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("CalculateScores() error = %v, wantErr %v", err, tt.wantErr)
			}
			for i, c := range tt.args.config.Task.Subtasks {
				if c.Score != tt.wantRes[i] {
					t.Errorf("CalculateScores() = %v, want %v", c.Score, tt.wantRes[i])
				}
			}
		})
	}
}

func TestCalculateScores(t *testing.T) {
	TestCalculateSimpleScores(t)
	TestCalculateSubtaskScores(t)
}
