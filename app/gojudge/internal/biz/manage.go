package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/uuid"
	pbc "sastoj/api/sastoj/gojudge/judger/gojudge/v1"
	pb "sastoj/api/sastoj/gojudge/judger/service/v1"
	"sastoj/app/gojudge/internal/data"
)

type ManageUseCase struct {
	goJudge []GoJudge
	logger  *log.Helper
	data    *data.Data
	cm      *CaseManage
}

func NewManageUseCase(logger log.Logger, data *data.Data, cm *CaseManage) *ManageUseCase {
	return &ManageUseCase{
		logger: log.NewHelper(logger),
		data:   data,
		cm:     cm,
	}
}

func (s *ManageUseCase) CreateOne(ctx context.Context, req *pb.CreateOneRequest) (*pb.CreateOneReply, error) {
	endpoint := req.Config.Endpoint
	ClientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
	)
	if err != nil {
		return nil, err
	}
	exec := pbc.NewExecutorClient(ClientConn)
	judge := GoJudge{
		Uuid: uuid.NewString(),
		Config: Config{
			Topic:    req.Config.Topic,
			Endpoint: req.Config.Endpoint,
		},
		cm:     s.cm,
		judge:  Judge{log: s.logger},
		db:     s.data,
		exec:   &exec,
		conn:   ClientConn,
		close:  nil,
		Status: StatusEnd,
	}
	err = judge.Start(ctx, judge.Config)
	if err != nil {
		return nil, err
	}
	s.goJudge = append(s.goJudge, judge)
	return &pb.CreateOneReply{
		Success: true,
		Status: &pb.Status{
			Uuid:   judge.Uuid,
			Status: judge.Status,
			Config: &pb.Config{
				Topic:    judge.Config.Topic,
				Endpoint: judge.Config.Endpoint,
			},
		},
	}, nil
}

func (s *ManageUseCase) FetchStatus(ctx context.Context, req *pb.FetchStatusRequest) (*pb.FetchStatusReply, error) {
	var target GoJudge
	if getUuid := req.GetUuid(); getUuid != "" {
		for _, judge := range s.goJudge {
			if judge.Uuid == getUuid {
				target = judge
				break
			}
		}
	}
	if &target == nil {
		return nil, errors.New("judge of uuid not found")
	}
	return &pb.FetchStatusReply{
		Status: &pb.Status{
			Uuid:   target.Uuid,
			Status: target.Status,
			Config: &pb.Config{
				Topic:    target.Config.Topic,
				Endpoint: target.Config.Endpoint,
			},
		},
	}, nil
}

func (s *ManageUseCase) ModifyConfig(ctx context.Context, req *pb.ModifyConfigRequest) (*pb.ModifyConfigReply, error) {
	//TODO config 重启单个judge
	return &pb.ModifyConfigReply{}, nil
}

func (s *ManageUseCase) DeleteOne(ctx context.Context, req *pb.DeleteOneRequest) (*pb.DeleteOneReply, error) {
	if getUuid := req.GetUuid(); getUuid != "" {
		for index, judge := range s.goJudge {
			if judge.Uuid == getUuid {
				s.goJudge = append(s.goJudge[:index], s.goJudge[index+1:]...)
				judge.Close()
				err := judge.conn.Close()
				if err != nil {
					return nil, err
				}
				break
			}
		}
	}
	return &pb.DeleteOneReply{
		Success: true,
	}, nil
}

func (s *ManageUseCase) ListAll(ctx context.Context, req *pb.ListAllRequest) (*pb.ListAllReply, error) {
	res := make([]*pb.Status, 0)
	for _, judge := range s.goJudge {
		res = append(res, &pb.Status{
			Uuid:   judge.Uuid,
			Status: judge.Status,
			Config: &pb.Config{
				Topic:    judge.Config.Topic,
				Endpoint: judge.Config.Endpoint,
			},
		})
	}
	return &pb.ListAllReply{
		Status: res,
	}, nil
}
