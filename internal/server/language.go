package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/internal/service"
	protos "github.com/iantal/ld/protos/ld"
)

type Linguist struct {
	log       hclog.Logger
	breakdown *service.Breakdown
}

func NewLinguist(l hclog.Logger) *Linguist {
	return &Linguist{l, service.NewBreakdown(l)}
}

func (l *Linguist) Breakdown(ctx context.Context, rr *protos.BreakdownRequest) (*protos.BreakdownResponse, error) {
	l.log.Info("Handle request for Breakdown", "projectID", rr.GetProjectID())

	breakdownResult, err := l.breakdown.GetLanguages(rr.GetProjectID())
	if err != nil {
		return nil, err
	}

	return &protos.BreakdownResponse{Breakdown: breakdownResult}, nil
}
