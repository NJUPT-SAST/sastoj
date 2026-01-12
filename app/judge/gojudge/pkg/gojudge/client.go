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

	cmd := &pb.Request_CmdType{}
	cmd.SetArgs(command.Compile)
	cmd.SetEnv(command.Env)
	cmd.SetFiles([]*pb.Request_File{
		requestMemory([]byte{}),
		requestPipe("stdout", command.CompileConfig.StdoutMaxSize),
		requestPipe("stderr", command.CompileConfig.StderrMaxSize),
	})
	cmd.SetCpuTimeLimit(command.CompileConfig.CpuTimeLimit * 1000 * 1000)
	cmd.SetCpuRateLimit(command.CompileConfig.CpuRateLimit)
	cmd.SetClockTimeLimit(command.CompileConfig.ClockTimeLimit * 1000 * 1000)
	cmd.SetMemoryLimit(command.CompileConfig.MemoryLimit * 1024 * 1024)
	cmd.SetProcLimit(command.CompileConfig.ProcLimit)
	cmd.SetCopyIn(map[string]*pb.Request_File{
		command.Source: requestMemory(code),
	})
	cmd.SetCopyOut([]*pb.Request_CmdCopyOutFile{
		requestCopyOutFile("stdout"),
		requestCopyOutFile("stderr"),
	})
	cmd.SetCopyOutCached([]*pb.Request_CmdCopyOutFile{
		requestCopyOutFile(command.Target),
	})

	req := &pb.Request{}
	req.SetRequestID(requestID)
	req.SetCmd([]*pb.Request_CmdType{cmd})

	res, err := (*g.Client).Exec(context.Background(), req)
	result, err := handleExecError(requestID, res, err)
	if err != nil {
		return "", result, err
	}
	return result.GetFileIDs()[command.Target], result, nil
}

func (g *GoJudge) ClassicJudge(problemId int64, inputStr string, input []byte, language string, targetID string, requestID string, cpuTimeLimit uint64, clockTimeLimit uint64, memoryLimit uint64, outputSize int64) (*pb.Response_Result, error) {
	//Tips: File should not be larger than Config
	command, ok := (*g.Commands)[language]
	if !ok {
		return nil, errors.New("language not support ")
	}
	var f []*pb.Request_File
	if problemId != -1 {
		f = []*pb.Request_File{
			requestLocal(problemId, inputStr),
			requestPipe("stdout", command.RunConfig.StdoutMaxSize),
			requestPipe("stderr", command.RunConfig.StderrMaxSize),
		}
	} else {
		f = []*pb.Request_File{
			requestMemory(input),
			requestPipe("stdout", command.RunConfig.StdoutMaxSize),
			requestPipe("stderr", command.RunConfig.StderrMaxSize),
		}
	}

	cmd := &pb.Request_CmdType{}
	cmd.SetArgs(command.Run)
	cmd.SetEnv([]string{"PATH=/usr/bin:/bin"})
	cmd.SetFiles(f)
	cmd.SetCpuTimeLimit(min(command.RunConfig.CpuTimeLimit, cpuTimeLimit) * 1000 * 1000)
	cmd.SetCpuRateLimit(command.RunConfig.CpuRateLimit)
	cmd.SetClockTimeLimit(min(command.RunConfig.ClockTimeLimit, clockTimeLimit) * 1000 * 1000)
	cmd.SetMemoryLimit(min(command.RunConfig.MemoryLimit, memoryLimit) * 1024 * 1024)
	cmd.SetProcLimit(command.RunConfig.ProcLimit)
	cmd.SetCopyIn(map[string]*pb.Request_File{
		command.Target: requestCached(targetID),
	})

	req := &pb.Request{}
	req.SetRequestID(requestID)
	req.SetCmd([]*pb.Request_CmdType{cmd})

	res, err := (*g.Client).Exec(context.Background(), req)
	result, err := handleExecError(requestID, res, err)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (g *GoJudge) AddFile(file string, content []byte) (fileID string, err error) {
	fc := &pb.FileContent{}
	fc.SetName(file)
	fc.SetContent(content)
	id, err := (*g.Client).FileAdd(context.Background(), fc)
	if err != nil {
		return "", err
	}
	return id.GetFileID(), nil
}

func (g *GoJudge) DeleteFile(fileID string) error {
	fid := &pb.FileID{}
	fid.SetFileID(fileID)
	_, err := (*g.Client).FileDelete(context.Background(), fid)
	return err
}

func handleExecError(requestID string, response *pb.Response, err error) (*pb.Response_Result, error) {
	if err != nil {
		return nil, err
	}
	results := response.GetResults()
	if len(results) == 0 {
		return nil, errors.New("request=" + requestID + ": no result")
	}
	result := results[0]
	if result.GetError() != "" {
		return result, errors.New("request=" + requestID + ": " + result.GetError())
	}
	if result.GetFileError() != nil {
		message := "request=" + requestID + ":\n"
		for _, fileError := range result.GetFileError() {
			message += "file error: " + fileError.GetMessage() + ",\n"
		}
		return result, errors.New(message)
	}
	if result.GetExitStatus() != 0 && result.GetExitStatus() < 32 {
		return result, errors.New("request=" + requestID + ": exit status=" + strconv.Itoa(int(result.GetExitStatus())))
	}
	return result, nil
}

func requestMemory(content []byte) *pb.Request_File {
	mem := &pb.Request_MemoryFile{}
	mem.SetContent(content)
	f := &pb.Request_File{}
	f.SetMemory(mem)
	return f
}

func requestLocal(problemId int64, inputStr string) *pb.Request_File {
	local := &pb.Request_LocalFile{}
	local.SetSrc("/data/cases/" + strconv.FormatInt(problemId, 10) + "/testdata/" + inputStr)
	f := &pb.Request_File{}
	f.SetLocal(local)
	return f
}

func requestPipe(name string, max int64) *pb.Request_File {
	pipe := &pb.Request_PipeCollector{}
	pipe.SetName(name)
	pipe.SetMax(max)
	f := &pb.Request_File{}
	f.SetPipe(pipe)
	return f
}

func requestCached(fileID string) *pb.Request_File {
	cached := &pb.Request_CachedFile{}
	cached.SetFileID(fileID)
	f := &pb.Request_File{}
	f.SetCached(cached)
	return f
}

func requestCopyOutFile(name string) *pb.Request_CmdCopyOutFile {
	cof := &pb.Request_CmdCopyOutFile{}
	cof.SetName(name)
	return cof
}
