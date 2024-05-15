package judge

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// GoJudge to
type GoJudge struct {
	//
	judge Judge

	//
	caseCacheManager CaseCacheManager

	log log.Helper

	ch *amqp.Channel

	loopGroups []*LoopGroup
}

type CaseCacheManager struct {
}

type LoopGroup struct {
	uuid   uuid.UUID
	closes []func()
}

func (group *LoopGroup) Close() {
	for _, c := range group.closes {
		c()
	}
	//TODO wait if passable
}

func (group *LoopGroup) Append(other *LoopGroup) {
	group.closes = append(group.closes, other.closes...)
}

func (g GoJudge) handleSubmit(ctx context.Context, v *Submit) error {
	return nil
}

func (g GoJudge) handleSelfTest(ctx context.Context, v *SelfTest) error {
	return nil
}

func (g GoJudge) StartOneLoop(ctx context.Context) (*LoopGroup, error) {
	closes := make([]func(), 0)
	closes = append(closes, g.judge.StartLoopOnSubmit(ctx, g.ch, g.handleSubmit))
	closes = append(closes, g.judge.StartLoopOnSelfTest(ctx, g.ch, g.handleSelfTest))
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	newLoopGroup := &LoopGroup{
		uuid:   newUUID,
		closes: closes,
	}
	g.loopGroups = append(g.loopGroups, newLoopGroup)
	return newLoopGroup, nil
}

func (g GoJudge) RefreshCaseCache(ctx context.Context) error { return nil }

func (g GoJudge) CloseAll() {
	for _, group := range g.loopGroups {
		for _, c := range group.closes {
			c()
		}
	}
	//TODO wait if passable
}
