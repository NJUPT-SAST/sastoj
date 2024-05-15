package biz

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	pb "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	"sastoj/app/gojudge/internal/data"
	"sastoj/ent"
	"sastoj/ent/problem"
	u "sastoj/pkg/util"
	"slices"
	"strconv"
)

const (
	// submit topic name
	submit = "submission"
	// selfTest topic name
	selfTest = "self-test"
	// StatusStart loop status
	StatusStart = "START"
	// StatusEnd loop status
	StatusEnd = "END"
)

// GoJudge is a Client for a single go-judge
type GoJudge struct {
	Uuid   string
	Status string
	Config Config
	judge  Judge
	db     *data.Data
	cm     *CaseManage
	exec   *pb.ExecutorClient
	conn   *grpc.ClientConn
	close  func()
}

type Config struct {
	Topic    string
	Endpoint string
}

// handleSubmit process the submission
func (g *GoJudge) handleSubmit(v *Submit) error {
	config, err := g.cm.GetConfig(v.ProblemID)
	if err != nil {
		return err
	}
	pro, err := g.db.Ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		First(context.Background())
	if err != nil {
		return err
	}
	fileID, err := compile(g.exec, v, config, "a", "")
	if err != nil {
		return err
	}
	totalTime := uint64(0)
	maxMemory := uint64(0)
	score := 0
	var builder []*ent.SubmissionCaseCreate
	for index, c := range config.Task.Cases {
		in, ans, err := g.cm.FetchCase(v.ProblemID, c.Input, c.Answer)
		if err != nil {
			_ = deleteFile(g.exec, fileID)
			return err
		}
		result, err := judge(g.exec, string(in), config, "a", fileID, "")
		if err != nil {
			_ = deleteFile(g.exec, fileID)
			return err
		}
		//non-Accept
		if result.Status.Number() != 1 || !slices.Equal(result.Files["stdout"], ans) {
			//get problem cases ID
			var problemCasesID int64
			for _, problemCase := range pro.Edges.ProblemCases {
				if problemCase.Index == int16(index) {
					problemCasesID = problemCase.ID
				}
			}
			create := g.db.Ent.SubmissionCase.Create().
				SetProblemCasesID(problemCasesID).
				SetMemory(int32(result.Memory)).
				SetTime(int32(result.RunTime)).
				SetMessage(result.Status.String()).
				SetState(int16(result.Status.Number())). //TODO conv to rs-judge state
				SetPoint(0)
			builder = append(builder, create)
			break
		}
		//Accept
		totalTime += result.RunTime
		if maxMemory < result.Memory {
			maxMemory = result.Memory
		}
		score += int(v.Point)
		//get problem cases ID
		var problemCasesID int64
		for _, problemCase := range pro.Edges.ProblemCases {
			if problemCase.Index == int16(index) {
				problemCasesID = problemCase.ID
			}
		}
		//gen builder
		create := g.db.Ent.SubmissionCase.Create().
			SetProblemCasesID(problemCasesID).
			SetMemory(int32(result.Memory)).
			SetTime(int32(result.RunTime)).
			SetMessage(result.Status.String()).
			SetState(int16(result.Status.Number())). //TODO conv to rs-judge state
			SetPoint(v.Point)
		builder = append(builder, create)
	}
	err = deleteFile(g.exec, fileID)
	if err != nil {
		return err
	}
	//Save into database
	cases, err := g.db.Ent.SubmissionCase.CreateBulk(builder...).Save(context.Background())
	if err != nil {
		return err
	}
	var ids []int64
	for _, ca := range cases {
		ids = append(ids, ca.ID)
	}
	err = g.db.Ent.Submission.Create().
		SetCaseVersion(int8(pro.CaseVersion)).
		SetUsersID(v.UserID).
		SetCode(v.Code).
		SetLanguage(v.Language).
		SetPoint(int16(score)).
		SetStatus(1). //TODO conv to rs-judge state
		SetCreateTime(v.CreateTime).
		SetTotalTime(int32(totalTime)).
		SetMaxMemory(int32(maxMemory)).
		SetProblemsID(v.ProblemID).
		AddSubmissionCaseIDs(ids...).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// handleSubmit process the self-test
func (g *GoJudge) handleSelfTest(v *SelfTest) error {
	//TODO self-test
	return nil
}

func memory(content string) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Memory{
			Memory: &pb.Request_MemoryFile{
				Content: []byte(content),
			},
		},
	}
}

func pipe(name string, max int64) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Pipe{
			Pipe: &pb.Request_PipeCollector{
				Name: name,
				Max:  max,
			},
		},
	}
}

func cached(fileID string) *pb.Request_File {
	return &pb.Request_File{
		File: &pb.Request_File_Cached{
			Cached: &pb.Request_CachedFile{
				FileID: fileID,
			},
		},
	}
}

func copyOutFile(name string) *pb.Request_CmdCopyOutFile {
	return &pb.Request_CmdCopyOutFile{
		Name: name,
	}
}

func compile(exec *pb.ExecutorClient, v *Submit, judgeConfig *u.JudgeConfig, file string, requestID string) (fileID string, err error) {
	res, err := (*exec).Exec(context.Background(), &pb.Request{
		RequestID: requestID,
		Cmd: []*pb.Request_CmdType{
			{
				//TODO config for other language
				Args: []string{"/usr/bin/g++", file + ".cc", "-o", file},
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*pb.Request_File{
					memory(""),
					pipe("stdout", 10240),
					pipe("stderr", 10240),
				},
				CpuTimeLimit: 10000000000,
				MemoryLimit:  104857600,
				ProcLimit:    50,
				CopyIn: map[string]*pb.Request_File{
					file + ".cc": memory(v.Code),
				},
				CopyOut: []*pb.Request_CmdCopyOutFile{
					copyOutFile("stdout"),
					copyOutFile("stderr"),
				},
				CopyOutCached: []*pb.Request_CmdCopyOutFile{
					copyOutFile(file),
				},
			},
		},
	})
	result, err := handleSingleResultResponseError(res, err)
	if err != nil {
		return
	}
	fileID = result.FileIDs[file]
	return fileID, nil
}

func judge(exec *pb.ExecutorClient, input string, judgeConfig *u.JudgeConfig, file string, fileID string, requestID string) (*pb.Response_Result, error) {
	res, err := (*exec).Exec(context.Background(), &pb.Request{
		RequestID: requestID,
		Cmd: []*pb.Request_CmdType{
			{
				Args: []string{file},
				Env:  []string{"PATH=/usr/bin:/bin"},
				Files: []*pb.Request_File{
					memory(input),
					pipe("stdout", 10240),
					pipe("stderr", 10240),
				},
				CpuTimeLimit:   uint64(judgeConfig.ResourceLimits.Time),
				ClockTimeLimit: uint64(judgeConfig.ResourceLimits.Time * 2),
				MemoryLimit:    uint64(judgeConfig.ResourceLimits.Memory),
				ProcLimit:      50,
				CopyIn: map[string]*pb.Request_File{
					file: cached(fileID),
				},
			},
		},
	})
	result, err := handleSingleResultResponseError(res, err)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func deleteFile(exec *pb.ExecutorClient, fileID string) error {
	_, err := (*exec).FileDelete(context.Background(), &pb.FileID{
		FileID: fileID,
	})
	return err
}

func handleSingleResultResponseError(response *pb.Response, err error) (*pb.Response_Result, error) {
	if err != nil {
		return nil, err
	}
	result := response.Results[0]
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	if result.FileError != nil {
		message := ""
		for _, fileError := range result.FileError {
			message += "File Error: " + fileError.Message + ", "
		}
		return nil, errors.New(message)
	}
	if result.ExitStatus != 0 {
		return nil, errors.New("Error Exit Status " + strconv.Itoa(int(result.ExitStatus)))
	}
	return result, nil
}

// Start add a loop to consume msg
func (g *GoJudge) Start(ctx context.Context, config Config) error {
	switch config.Topic {
	case selfTest:
		{
			if g.close == nil || g.Status == StatusStart {
				c2 := g.judge.StartLoopOnSelfTest(g.db.Ch, selfTest, g.handleSelfTest)
				g.close = c2
				g.Status = StatusStart
			} else {
				return errors.New("has been started")
			}
		}
	case submit:
		{
			if g.close == nil || g.Status == StatusStart {
				c1 := g.judge.StartLoopOnSubmit(g.db.Ch, submit, g.handleSubmit)
				g.close = c1
				g.Status = StatusStart
			} else {
				return errors.New("has been started")
			}
		}
	default:
		return errors.New("invalid topic")
	}
	return nil
}

func (g *GoJudge) Close() {
	g.close()
	g.close = nil
	g.Status = StatusEnd
}
