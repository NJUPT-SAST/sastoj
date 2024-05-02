package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
	v1 "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/biz"
	"sastoj/app/user/gateway/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewContestRepo, NewProblemRepo, NewSubmissionRepo, NewContestServiceClient)

// Data .
type Data struct {
	cache *Cache
	ch    *amqp.Channel
	cc    v1.ContestServiceClient
}

// Cache is a map store.
type Cache struct {
	group2contests   map[int64][]*biz.Contest
	contest2problems map[int64][]*biz.Problem
	problems         map[int64]*biz.Problem
	submissions      map[string]*biz.Submission
}

// NewData .
func NewData(c *conf.Data, cc v1.ContestServiceClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// connect to redis
	cache := &Cache{
		group2contests:   make(map[int64][]*biz.Contest),
		contest2problems: make(map[int64][]*biz.Problem),
		problems:         make(map[int64]*biz.Problem),
		submissions:      make(map[string]*biz.Submission),
	}
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
		cache.group2contests[groupID] = contests
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
			log.Errorf("failed getting contest2problems: %v", err)
			return nil, nil, err
		}
		for _, problem := range problemsDTO.Problems {
			problems = append(problems, &biz.Problem{
				ID:    problem.Id,
				Title: problem.Title,
				Point: int16(problem.Point),
				Index: int16(problem.Index),
			})
			problemDTO, err := cc.GetProblem(context.Background(), &v1.GetProblemRequest{
				ProblemId: problem.Id,
			})
			if err != nil {
				log.Errorf("failed getting problem: %v", err)
			}
			cache.problems[problem.Id] = &biz.Problem{
				ID:      problem.Id,
				Title:   problem.Title,
				Content: problemDTO.Content,
				Point:   int16(problem.Point),
			}
		}
		cache.contest2problems[contest.Id] = problems
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
		cache: cache,
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
