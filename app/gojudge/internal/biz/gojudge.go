package biz

import (
	"context"
	"errors"
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
)

// Gojudge is a Client for a single go-judge
type Gojudge struct {
	Endpoint string
	Db       *data.Data
	close    []func()
}

// handleSubmit process the submission
func (g *Gojudge) handleSubmit(v *Submit) error {
	client := g.Db.Clients[g.Endpoint]
	config, err := g.Db.CaseManage.GetConfig(v.ProblemID)
	if err != nil {
		return err
	}
	pro, err := g.Db.Ent.Problem.Query().
		Where(problem.ID(v.ProblemID)).
		First(context.Background())
	if err != nil {
		return err
	}
	fileID, err := compile(client, v, config, "a", "")
	if err != nil {
		return err
	}
	totalTime := uint64(0)
	maxMemory := uint64(0)
	score := 0
	var builder []*ent.SubmissionCaseCreate
	for index, c := range config.Task.Cases {
		in, ans, err := g.Db.CaseManage.FetchCase(v.ProblemID, c.Input, c.Answer)
		if err != nil {
			_ = deleteFile(client, fileID)
			return err
		}
		result, err := judge(client, string(in), config, "a", fileID, "")
		if err != nil {
			_ = deleteFile(client, fileID)
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
			create := g.Db.Ent.SubmissionCase.Create().
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
		create := g.Db.Ent.SubmissionCase.Create().
			SetProblemCasesID(problemCasesID).
			SetMemory(int32(result.Memory)).
			SetTime(int32(result.RunTime)).
			SetMessage(result.Status.String()).
			SetState(int16(result.Status.Number())). //TODO conv to rs-judge state
			SetPoint(v.Point)
		builder = append(builder, create)
	}
	err = deleteFile(client, fileID)
	if err != nil {
		return err
	}
	//Save into database
	cases, err := g.Db.Ent.SubmissionCase.CreateBulk(builder...).Save(context.Background())
	if err != nil {
		return err
	}
	var ids []int64
	for _, ca := range cases {
		ids = append(ids, ca.ID)
	}
	err = g.Db.Ent.Submission.Create().
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
func (g *Gojudge) handleSelfTest(v *SelfTest) error {
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
			message += "file error: " + fileError.Message + ", "
		}
		return nil, errors.New(message)
	}
	if result.ExitStatus != 0 {
		return nil, errors.New("error exit status " + strconv.Itoa(int(result.ExitStatus)))
	}
	return result, nil
}

// Start add a loop to consume msg
func (g *Gojudge) Start() error {
	if g.close == nil {
		c1 := StartLoopOnSubmit(g.Db.Logger, g.Db.Ch, submit, g.handleSubmit)
		c2 := StartLoopOnSelfTest(g.Db.Logger, g.Db.Ch, selfTest, g.handleSelfTest)
		g.close = []func(){c1, c2}
	} else {
		return errors.New("loops on-submit, on-self-test has been started")
	}
	return nil
}

func (g *Gojudge) Close() {
	for _, c := range g.close {
		c()
	}
	g.close = nil
}
