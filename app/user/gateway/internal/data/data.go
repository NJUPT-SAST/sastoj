package data

import (
	"context"
	"errors"
	"fmt"
	"net"
	v1 "sastoj/api/sastoj/user/contest/service/v1"
	"sastoj/app/user/gateway/internal/biz"
	"sastoj/app/user/gateway/internal/conf"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	amqp "github.com/rabbitmq/amqp091-go"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewContestRepo, NewProblemRepo, NewSubmissionRepo, NewContestClient)

// Data .
type Data struct {
	cache *Cache
	ch    *amqp.Channel
	cc    v1.ContestClient
}

// Cache is a map store.
type Cache struct {
	contestantsMap   map[int64][]*biz.Contest
	contest2problems map[int64][]*biz.Problem
	problems         map[int64]*biz.Problem
	submissions      map[string]*biz.Submission
	selfTests        map[string]*biz.SelfTest
	cases            map[string][]*biz.Case
	token            string
}

// NewData .
func NewData(c *conf.Data, cc v1.ContestClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// connect to redis
	cache := &Cache{
		contestantsMap:   make(map[int64][]*biz.Contest),
		contest2problems: make(map[int64][]*biz.Problem),
		problems:         make(map[int64]*biz.Problem),
		submissions:      make(map[string]*biz.Submission),
		selfTests:        make(map[string]*biz.SelfTest),
		cases:            make(map[string][]*biz.Case),
	}
	// get registerGateway IPv4
	ip, err := getIPv4(byte(c.Ipv4Prefix))
	if err != nil {
		return nil, nil, err
	}
	endpoint := ip + ":" + strconv.FormatInt(c.Port, 10)
	registerGateway, err := cc.Register(context.Background(), &v1.RegisterRequest{
		Endpoint: endpoint,
		Secret:   c.Secret,
	})
	if err != nil {
		return nil, nil, err
	}
	cache.token = registerGateway.Token
	// get today contests
	contestsDTO, err := cc.GetContests(context.Background(), &v1.GetContestsRequest{})
	if err != nil {
		return nil, nil, err
	}
	contestantsMap := make(map[int64][]*biz.Contest)
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
		for _, groupID := range contestDTO.Contestants {
			contestantsMap[groupID] = append(contestantsMap[groupID], contest)
		}
	}
	for groupID, contests := range contestantsMap {
		cache.contestantsMap[groupID] = contests
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
				Type:  problem.Type,
				Score: problem.Score,
				Index: int16(problem.Index),
			})
			problemDTO, err := cc.GetProblem(context.Background(), &v1.GetProblemRequest{
				ProblemId: problem.Id,
			})
			if err != nil {
				log.Errorf("failed getting problem: %v", err)
			}
			cache.problems[problem.Id] = &biz.Problem{
				ID:       problem.Id,
				Title:    problem.Title,
				Type:     problem.Type,
				Content:  problemDTO.Content,
				Score:    problem.Score,
				Index:    int16(problem.Index),
				Metadata: problemDTO.Metadata,
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

func NewContestClient(conf *conf.Client) v1.ContestClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(conf.Grpc.Addr),
	)
	if err != nil {
		panic(err)
	}
	cc := v1.NewContestClient(conn)
	return cc
}

func getIPv4(ipv4Prefix byte) (string, error) {
	// Get the current client's IPs
	adds, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return "", errors.New("system IPv4 not found")
	}
	for _, address := range adds {
		if aspnet, ok := address.(*net.IPNet); !(!ok || aspnet.IP.IsLoopback()) {
			if aspnet.IP.To4() != nil && aspnet.IP.To4()[0] == ipv4Prefix {
				return aspnet.IP.String(), nil
			}
		}
	}
	return "", fmt.Errorf("get %s, but IPv4 start with %d not found", adds, ipv4Prefix)
	//return "", errors.New("IPv4 start with " + strconv.Itoa(int(ipv4Prefix)) + " not found")
}
