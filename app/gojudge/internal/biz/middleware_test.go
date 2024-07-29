package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"
	pb "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/app/gojudge/internal/conf"
	"sastoj/app/gojudge/internal/data"
	"testing"
	"time"
)

func TestHandleSubmit(t *testing.T) {

	client, err := data.GenEnt("postgres", "host=localhost port=5432 user=postgres dbname=sastoj password=123456789 sslmode=disable", log.GetLogger())

	language := "Bash"
	code := "#!/bin/bash\n\n# Function to perform sorting\nsort_numbers() {\necho \"Enter numbers separated by spaces:\"\nread -a numbers\n\nsorted_numbers=($(for num in \"${numbers[@]}\"; do echo $num; done | sort -n))\n\necho \"Sorted numbers:\"\necho \"${sorted_numbers[@]}\"\n}\n\n# Main script execution\nsort_numbers"

	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))
	exec := pb.NewExecutorClient(ClientConn)

	commands := data.NewCommands(
		[]string{"Bash"},
		map[string]string{},
		map[string]string{
			"Bash": "/bin/bash foo.sh",
		},
		map[string]string{},
		map[string]string{
			"Bash": "foo.sh",
		},
		map[string]*conf.LanguageConfig{
			"default": {
				Compile: &conf.ExecConfig{
					ProcLimit:      50,
					CpuTimeLimit:   10000000000,
					CpuRateLimit:   10000000000,
					ClockTimeLimit: 100000000000,
					MemoryLimit:    104857600,
					StdoutMaxSize:  100000000,
					StderrMaxSize:  100000000,
				},
				Run: &conf.ExecConfig{
					ProcLimit:      50,
					CpuTimeLimit:   10000000000,
					CpuRateLimit:   10000000000,
					ClockTimeLimit: 100000000000,
					MemoryLimit:    104857600,
					StdoutMaxSize:  100000000,
					StderrMaxSize:  100000000,
				}},
		})

	goJudge := GoJudge{
		client:   &exec,
		commands: &commands,
	}
	if err != nil {
		panic(err)
	}

	middleware := &Middleware{
		ent:        client,
		redis:      nil,
		fileManage: &data.FileManage{FileLocation: "C:\\Users\\Dell\\Desktop\\load"},
		goJudge:    &goJudge,
		logger:     log.NewHelper(log.With(log.GetLogger(), "module", "judge-middleware")),
		close:      nil,
	}

	err = middleware.handleSubmit(&Submit{
		ID:         uuid.NewString(),
		UserID:     1,
		ProblemID:  100,
		Code:       code,
		Status:     0,
		Point:      100,
		CreateTime: time.Now(),
		TotalTime:  0,
		MaxMemory:  0,
		Language:   language,
		Token:      "",
	})

	if err != nil {
		log.Infof("handle error:  %v", err)
	}

}
