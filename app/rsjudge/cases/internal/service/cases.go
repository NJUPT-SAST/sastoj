package service

import (
	"bytes"
	"google.golang.org/grpc/metadata"
	"io"
	pb "sastoj/api/sastoj/rsjudge/cases/service/v1"
	"sastoj/app/rsjudge/cases/internal/biz"
	"strconv"
)

type CasesService struct {
	pb.UnimplementedCasesServiceServer
	uc *biz.CasesUsecase
}

func NewCasesService(usecase *biz.CasesUsecase) *CasesService {
	return &CasesService{
		uc: usecase,
	}
}

func (s *CasesService) FetchCases(req *pb.FetchCasesRequest, conn pb.CasesService_FetchCasesServer) error {
	testdata, chunkSize, err := s.uc.FetchCases(int64(req.ProblemId))
	if err != nil {
		return err
	}
	err = conn.SendHeader(metadata.New(map[string]string{
		"file-name": strconv.FormatInt(int64(req.ProblemId), 10),
		"file-type": "tar.zst",
		"file-size": strconv.Itoa(len(testdata)),
	}))
	if err != nil {
		return err
	}
	chunk := &pb.FetchCasesResponse{Chunk: make([]byte, chunkSize)}
	var n int
	reader := bytes.NewReader(testdata)
	for {
		n, err = reader.Read(chunk.Chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk.Chunk = chunk.Chunk[:n]
		err = conn.Send(chunk)
		if err != nil {
			return err
		}
	}
	return nil
}
