package gojudge

import (
	"context"
	"errors"
	"strconv"

	"github.com/criyle/go-judge/pb"
)

type GoJudge struct {
	Client   *pb.ExecutorClient
	Commands *Commands
}

func (g *GoJudge) CheckLanguage(language string) error {
	if _, ok := (*g.Commands)[language]; !ok {
		return errors.New("language not support on this judge-mid")
	}
	return nil
}

func (g *GoJudge) Compile(code []byte, language string, requestID string) (string, *pb.Response_Result, error) {
	command, ok := (*g.Commands)[language]
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
	res, err := (*g.Client).Exec(context.Background(), &pb.Request{
		RequestID: requestID,
		Cmd: []*pb.Request_CmdType{
			{
				Args: command.Compile,
				Env:  command.Env,
				Files: []*pb.Request_File{
					requestMemory([]byte{}),
					requestPipe("stdout", command.CompileConfig.StdoutMaxSize),
					requestPipe("stderr", command.CompileConfig.StderrMaxSize),
				},
				CpuTimeLimit:   command.CompileConfig.CpuTimeLimit * 1000 * 1000,
				CpuRateLimit:   command.CompileConfig.CpuRateLimit,
				ClockTimeLimit: command.CompileConfig.ClockTimeLimit * 1000 * 1000,
				MemoryLimit:    command.CompileConfig.MemoryLimit * 1024 * 1024,
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
	result, err := handleExecError(requestID, res, err)
	if err != nil {
		return "", result, err
	}
	return result.FileIDs[command.Target], result, nil
}

func (g *GoJudge) ClassicJudge(input []byte, language string, targetID string, requestID string, cpuTimeLimit uint64, clockTimeLimit uint64, memoryLimit uint64, outputSize int64) (*pb.Response_Result, error) {
	//Tips: File should not be larger than Config
	command, ok := (*g.Commands)[language]
	if !ok {
		return nil, errors.New("language not support ")
	}
	res, err := (*g.Client).Exec(context.Background(), &pb.Request{
		RequestID: requestID,
		Cmd: []*pb.Request_CmdType{
			{
				Args: command.Run,
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*pb.Request_File{
					requestMemory(input),
					requestPipe("stdout", command.RunConfig.StdoutMaxSize),
					requestPipe("stderr", command.RunConfig.StderrMaxSize),
				},
				CpuTimeLimit:   min(command.RunConfig.CpuTimeLimit, cpuTimeLimit) * 1000 * 1000,
				CpuRateLimit:   command.RunConfig.CpuRateLimit,
				ClockTimeLimit: min(command.RunConfig.ClockTimeLimit, clockTimeLimit) * 1000 * 1000,
				MemoryLimit:    min(command.RunConfig.MemoryLimit, memoryLimit) * 1024 * 1024,
				ProcLimit:      command.RunConfig.ProcLimit,
				CopyIn: map[string]*pb.Request_File{
					command.Target: requestCached(targetID),
				},
			},
		},
	})
	result, err := handleExecError(requestID, res, err)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (g *GoJudge) AddFile(file string, content []byte) (fileID string, err error) {
	id, err := (*g.Client).FileAdd(context.Background(), &pb.FileContent{
		Name:    file,
		Content: content,
	})
	if err != nil {
		return "", err
	}
	return id.FileID, nil
}

func (g *GoJudge) DeleteFile(fileID string) error {
	_, err := (*g.Client).FileDelete(context.Background(), &pb.FileID{
		FileID: fileID,
	})
	return err
}

func handleExecError(requestID string, response *pb.Response, err error) (*pb.Response_Result, error) {
	if err != nil {
		return nil, err
	}
	if len(response.Results) == 0 {
		return nil, errors.New("request=" + requestID + ": no result")
	}
	result := response.Results[0]
	if result.Error != "" {
		return result, errors.New("request=" + requestID + ": " + result.Error)
	}
	if result.FileError != nil {
		message := "request=" + requestID + ":\n"
		for _, fileError := range result.FileError {
			message += "file error: " + fileError.Message + ",\n"
		}
		return result, errors.New(message)
	}
	if result.ExitStatus != 0 && result.ExitStatus < 32 {
		return result, errors.New("request=" + requestID + ": exit status=" + strconv.Itoa(int(result.ExitStatus)))
	}
	return result, nil
}

func requestMemory(content []byte) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Memory{
			Memory: &pb.Request_MemoryFile{
				Content: content,
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
