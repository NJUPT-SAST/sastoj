package biz

import (
	"context"
	pb "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/app/gojudge/internal/conf"
	"sastoj/app/gojudge/internal/data"
	"sastoj/pkg/util"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// TestHandleSubmit require start with the env: go-judge(diff languages)
func TestCGoJudge(t *testing.T) {
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
		map[string][]string{
			"default": {"PATH=/usr/bin:/bin"},
		},
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
	fileID, res, err := goJudge.Compile([]byte(code), language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge([]byte(input), language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
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

// TestHandleSubmit require start with the env: go-judge(diff languages)
func TestBashGoJudge(t *testing.T) {
	language := "Bash"
	code := `#!/bin/bash

	read a b
	sum=$((a + b))

	echo "$sum"`
	input := util.RemoveCr([]byte{51, 49, 48, 55, 53, 32, 50, 50, 52, 51, 48, 13, 10})
	inputs := []string{string(input), "121312 512312\n"}
	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))
	exec := pb.NewExecutorClient(ClientConn)

	commands := data.NewCommands(
		[]string{"Bash"},
		map[string][]string{
			"default": {"PATH=/usr/bin:/bin"},
		},
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

	//compile
	fileID, res, err := goJudge.Compile([]byte(code), language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge([]byte(input), language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
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

// TODO: fix Java build
func TestJavaGoJudge(t *testing.T) {
	language := "Java"
	code := `import java.util.Scanner;

public class Main {
	public static void main(String[] args) {
		Scanner sc = new Scanner(System.in);
		int a = sc.nextInt();
		int b = sc.nextInt();
		System.out.println(a + b);
	}
}`
	inputs := []string{"1 2\n", "3 4\n"}
	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))
	if err != nil {
		panic(err)
	}
	exec := pb.NewExecutorClient(ClientConn)

	enable := []string{"Java"}
	env := map[string][]string{
		"default": {"PATH=/usr/bin:/bin"},
	}
	compile := map[string]string{
		"Java": `/usr/bin/bash -c "/usr/bin/javac Main.java && /usr/bin/jar cvf Main.jar *.class > /dev/null"`,
	}
	run := map[string]string{
		"Java": "/usr/bin/java Main.jar",
	}
	source := map[string]string{
		"Java": "Main.java",
	}
	target := map[string]string{
		"Java": "Main.jar",
	}
	configs := map[string]*conf.LanguageConfig{
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
			},
		},
	}
	command := data.NewCommands(
		enable,
		env,
		compile,
		run,
		source,
		target,
		configs,
	)

	goJudge := GoJudge{
		client:   &exec,
		commands: &command,
	}

	//compile
	fileID, res, err := goJudge.Compile([]byte(code), language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge([]byte(input), language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
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

func TestGolangGoJudge(t *testing.T) {
	language := "Golang"
	code := `package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println(a + b)
}`
	inputs := []string{"1 2\n", "3 4\n"}
	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))
	if err != nil {
		panic(err)
	}
	exec := pb.NewExecutorClient(ClientConn)

	enable := []string{"Golang"}
	env := map[string][]string{
		"Golang": {"PATH=/usr/bin:/bin", "GOCACHE=/tmp/"},
	}
	compile := map[string]string{
		"Golang": "/usr/bin/go build -o foo foo.go",
	}
	run := map[string]string{
		"default": "foo",
	}
	source := map[string]string{
		"Golang": "foo.go",
	}
	target := map[string]string{
		"Golang": "foo",
	}
	configs := map[string]*conf.LanguageConfig{
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
			},
		},
	}
	command := data.NewCommands(
		enable,
		env,
		compile,
		run,
		source,
		target,
		configs,
	)

	goJudge := GoJudge{
		client:   &exec,
		commands: &command,
	}

	//compile
	fileID, res, err := goJudge.Compile([]byte(code), language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge([]byte(input), language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
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

func TestPython3GoJudge(t *testing.T) {
	language := "Python"
	code := `a, b = map(int, input().split())
print(a + b)`
	inputs := []string{"1 2\n", "3 4\n"}
	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))
	if err != nil {
		panic(err)
	}
	exec := pb.NewExecutorClient(ClientConn)

	enable := []string{"Python"}
	env := map[string][]string{
		"Python": {"PATH=/usr/bin:/bin"},
	}
	compile := map[string]string{
		"Python": "",
	}
	run := map[string]string{
		"Python": "/usr/bin/python3 foo.py",
	}
	source := map[string]string{
		"Python": "foo.py",
	}
	target := map[string]string{
		"Python": "foo.py",
	}
	configs := map[string]*conf.LanguageConfig{
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
			},
		},
	}
	command := data.NewCommands(
		enable,
		env,
		compile,
		run,
		source,
		target,
		configs,
	)

	goJudge := GoJudge{
		client:   &exec,
		commands: &command,
	}

	//compile
	fileID, res, err := goJudge.Compile([]byte(code), language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge([]byte(input), language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
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

func TestPHPGoJudge(t *testing.T) {
	language := "PHP"
	code := `<?php
$line = fgets(STDIN);
$nums = explode(" ", $line);
echo $nums[0] + $nums[1];`
	inputs := []string{"1 2\n", "3 4\n"}
	endpoint := "127.0.0.1:5051"

	//connect to go-judge
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithHealthCheck(false))

	if err != nil {
		panic(err)
	}

	exec := pb.NewExecutorClient(ClientConn)

	enable := []string{"PHP"}
	env := map[string][]string{
		"PHP": {"PATH=/usr/bin:/bin"},
	}
	compile := map[string]string{
		"PHP": "",
	}
	run := map[string]string{
		"PHP": "/usr/bin/php foo.php",
	}
	source := map[string]string{
		"PHP": "foo.php",
	}
	target := map[string]string{
		"PHP": "foo.php",
	}
	configs := map[string]*conf.LanguageConfig{
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
			},
		},
	}
	command := data.NewCommands(
		enable,
		env,
		compile,
		run,
		source,
		target,
		configs,
	)

	goJudge := GoJudge{
		client:   &exec,
		commands: &command,
	}

	//compile
	fileID, res, err := goJudge.Compile([]byte(code), language, "1")
	if err != nil {
		log.Errorf("failed compile file: %v  \nres :%v", err, res)
		panic(err)
	}
	log.Infof("compiled fileID: %v", fileID)

	//judge
	for _, input := range inputs {
		judge, err := goJudge.Judge([]byte(input), language, fileID, "2", 10000000000, 10000000000*2, 104857600, 1240000)
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
