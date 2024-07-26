package biz

import (
	"context"
	"errors"
	pb "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/app/gojudge/internal/data"
	"strconv"
)

type GoJudge struct {
	client   *pb.ExecutorClient
	commands *data.Commands
}

func (g *GoJudge) CheckLanguage(language string) error {
	if _, ok := (*g.commands)[language]; !ok {
		return errors.New("language not support on this judge")
	}
	return nil
}

func (g *GoJudge) Compile(code string, language string, requestID string) (string, *pb.Response_Result, error) {
	command, ok := (*g.commands)[language]
	if !ok {
		return "", nil, errors.New("language not support ")
	}
	if len(command.Compile) == 0 {
		fileID, err := g.AddFile(command.Target, code)
		if err != nil {
			return "", nil, err
		}
		return fileID, nil, nil
	}
	res, err := (*g.client).Exec(context.Background(), &pb.Request{
		RequestID: requestID,
		Cmd: []*pb.Request_CmdType{
			{
				Args: command.Compile,
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*pb.Request_File{
					requestMemory(""),
					requestPipe("stdout", command.CompileConfig.StdoutMaxSize),
					requestPipe("stderr", command.CompileConfig.StderrMaxSize),
				},
				CpuTimeLimit:   command.CompileConfig.CpuTimeLimit,
				CpuRateLimit:   command.CompileConfig.CpuRateLimit,
				ClockTimeLimit: command.CompileConfig.ClockTimeLimit,
				MemoryLimit:    command.CompileConfig.MemoryLimit,
				ProcLimit:      command.CompileConfig.ProcLimit,
				CopyIn: map[string]*pb.Request_File{
					command.Source: requestMemory(code),
				},
				CopyOut: []*pb.Request_CmdCopyOutFile{
					requestCopyOutFile("stdout"),
					requestCopyOutFile("stderr"),
				},
				CopyOutCached: []*pb.Request_CmdCopyOutFile{
					requestCopyOutFile(command.Target),
				},
			},
		},
	})
	result, err := handleExecError(res, err)
	if err != nil {
		return "", result, err
	}
	return result.FileIDs[command.Target], result, nil
}

func (g *GoJudge) Judge(input string, language string, targetID string, requestID string, cpuTimeLimit uint64, clockTimeLimit uint64, memoryLimit uint64, outputSize int64) (*pb.Response_Result, error) {
	//Tips: File should not be larger than Config
	command, ok := (*g.commands)[language]
	if !ok {
		return nil, errors.New("language not support ")
	}
	res, err := (*g.client).Exec(context.Background(), &pb.Request{
		RequestID: requestID,
		Cmd: []*pb.Request_CmdType{
			{
				Args: command.Run,
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*pb.Request_File{
					requestMemory(input),
					requestPipe("stdout", min(command.RunConfig.StdoutMaxSize, outputSize)),
					requestPipe("stderr", command.RunConfig.StderrMaxSize),
				},
				CpuTimeLimit:   min(command.RunConfig.CpuTimeLimit, cpuTimeLimit),
				CpuRateLimit:   command.RunConfig.CpuRateLimit,
				ClockTimeLimit: min(command.RunConfig.ClockTimeLimit, clockTimeLimit),
				MemoryLimit:    min(command.RunConfig.MemoryLimit, memoryLimit),
				ProcLimit:      command.RunConfig.ProcLimit,
				CopyIn: map[string]*pb.Request_File{
					command.Target: requestCached(targetID),
				},
			},
		},
	})
	result, err := handleExecError(res, err)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (g *GoJudge) AddFile(file string, content string) (fileID string, err error) {
	id, err := (*g.client).FileAdd(context.Background(), &pb.FileContent{
		Name:    file,
		Content: []byte(content),
	})
	if err != nil {
		return "", err
	}
	return id.FileID, nil
}

func (g *GoJudge) DeleteFile(fileID string) error {
	_, err := (*g.client).FileDelete(context.Background(), &pb.FileID{
		FileID: fileID,
	})
	return err
}

func handleExecError(response *pb.Response, err error) (*pb.Response_Result, error) {
	if err != nil {
		return nil, err
	}
	result := response.Results[0]
	if result.Error != "" {
		return result, errors.New(result.Error)
	}
	if result.FileError != nil {
		message := ""
		for _, fileError := range result.FileError {
			message += "file error: " + fileError.Message + ", "
		}
		return result, errors.New(message)
	}
	if result.ExitStatus != 0 {
		return result, errors.New("error exit status " + strconv.Itoa(int(result.ExitStatus)))
	}
	return result, nil
}

func requestMemory(content string) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Memory{
			Memory: &pb.Request_MemoryFile{
				Content: []byte(content),
			},
		},
	}
}

func requestPipe(name string, max int64) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Pipe{
			Pipe: &pb.Request_PipeCollector{
				Name: name,
				Max:  max,
			},
		},
	}
}

func requestCached(fileID string) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Cached{
			Cached: &pb.Request_CachedFile{
				FileID: fileID,
			},
		},
	}
}

func requestCopyOutFile(name string) *pb.Request_CmdCopyOutFile {
	return &pb.Request_CmdCopyOutFile{
		Name: name,
	}
}
