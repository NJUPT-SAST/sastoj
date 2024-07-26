package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	pb "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/app/gojudge/internal/conf"
	"sastoj/app/gojudge/internal/data"
	"testing"
)

// TestHandleSubmit require start with the env: go-judge(diff languages)
func TestGoJudge(t *testing.T) {
	language := "C++"
	code := "#include <iostream>\n#include <vector>\n#include <algorithm>\n int main(){int n;std::cin >> n" +
		";std::vector<int> numbers(n);for(int i = 0; i < n; ++i) {std::cin >> numbers[i];}std::sort(numbers." +
		"begin(), numbers.end());for(const int& num : numbers){std::cout << num << \" \";}std::cout << std::endl;" +
		"return 0;}"
	inputs := []string{"5\n4 1 3 2 5"}
	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))
	exec := pb.NewExecutorClient(ClientConn)

	commands := data.NewCommands(
		[]string{"C", "C++"},
		map[string]string{
			"C":   "/usr/bin/gcc -Wall --std=c99 -o foo foo.c -lm",
			"C++": "/usr/bin/g++ -Wall -std=c++14 -o foo foo.cc -lm -I/include",
		},
		map[string]string{
			"default": "foo",
		},
		map[string]string{
			"C":   "foo.c",
			"C++": "foo.cc",
		},
		map[string]string{
			"default": "foo",
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

	//compile
	fileID, res, err := goJudge.Compile(code, language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge(input, language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
		if err != nil {
			log.Errorf("failed running judge: %v", err)
			panic(err)
		}
		log.Infof("judge: %v", judge)
	}

	//delete compiled file
	err = goJudge.DeleteFile(fileID)
	if err != nil {
		log.Errorf("failed deleting file: %v", err)
		panic(err)
	}
}
