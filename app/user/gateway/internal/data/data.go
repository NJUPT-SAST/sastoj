package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	v1 "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/biz"
	"sastoj/app/user/gateway/internal/conf"
	"strconv"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewContestRepo, NewProblemRepo, NewSubmissionRepo, NewContestServiceClient)

// Data .
type Data struct {
	redis *redis.Client
	ch    *amqp.Channel
	cc    v1.ContestServiceClient
}

// NewData .
func NewData(c *conf.Data, cc v1.ContestServiceClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// connect to redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		DB:   int(c.Redis.Db),
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Errorf("failed connecting to redis: %v", err)
		return nil, nil, err
	}

	// init redis
	// get today contests
	contestsDTO, err := cc.GetContests(context.Background(), &v1.GetContestsRequest{})
	if err != nil {
		return nil, nil, err
	}
	groupID2contests := make(map[int64][]*biz.Contest)
	for _, contestDTO := range contestsDTO.Contests {
		contest := &biz.Contest{
			ID:          contestDTO.Id,
			Title:       contestDTO.Title,
			Description: contestDTO.Description,
			Status:      int16(contestDTO.State),
			Type:        int16(contestDTO.Type),
			StartTime:   contestDTO.StartTime.AsTime(),
			EndTime:     contestDTO.EndTime.AsTime(),
			Language:    contestDTO.Language,
			ExtraTime:   int16(contestDTO.ExtraTime),
		}
		for _, groupID := range contestDTO.Groups {
			groupID2contests[groupID] = append(groupID2contests[groupID], contest)
		}
	}
	for groupID, contests := range groupID2contests {
		redisClient.Set(context.Background(), "group:"+strconv.FormatInt(groupID, 10), contests, 24*time.Hour)
	}

	// get problems
	var problems []*biz.Problem
	for _, contestDTO := range contestsDTO.Contests {
		contest := contestDTO
		// go func() {
		problemsDTO, err := cc.GetProblems(context.Background(), &v1.GetProblemsRequest{
			ContestId: contest.Id,
		})
		if err != nil {
			log.Errorf("failed getting problems: %v", err)
			return nil, nil, err
		}
		for _, problemDTO := range problemsDTO.Problems {
			problems = append(problems, &biz.Problem{
				ID:    problemDTO.Id,
				Title: problemDTO.Title,
				Point: int16(problemDTO.Point),
				Index: int16(problemDTO.Index),
			})
		}
		redisClient.Set(context.Background(), "contest:"+strconv.FormatInt(contest.Id, 10), problems, 24*time.Hour)
		// }()
	}

	// get problem
	for _, problem := range problems {

		problemID := problem.ID
		// go func() {
		problemDTO, err := cc.GetProblem(context.Background(), &v1.GetProblemRequest{
			ProblemId: problemID,
		})
		if err != nil {
			return nil, nil, err
		}
		problem := &biz.Problem{
			ID:      problemDTO.Id,
			Title:   problemDTO.Title,
			Content: problemDTO.Content,
			Point:   int16(problemDTO.Point),
		}
		redisClient.Set(context.Background(), "problem:"+strconv.FormatInt(problem.ID, 10), problem, 24*time.Hour)
		// }()
	}

	// connect to mq
	conn, err := amqp.Dial(c.Mq)
	if err != nil {
		log.Errorf("failed connecting to mq: %v", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Errorf("failed opening a channel")
	}
	return &Data{
		redis: redisClient,
		cc:    cc,
		ch:    ch,
	}, cleanup, nil
}

func NewContestServiceClient(conf *conf.Client) v1.ContestServiceClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Grpc.Endpoint),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewContestServiceClient(conn)
	return c
}
