package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/internal/files"
	"github.com/iantal/ld/internal/service"
	protos "github.com/iantal/ld/protos/ld"
)

type Linguist struct {
	log       hclog.Logger
	breakdown *service.Breakdown
}

func NewLinguist(l hclog.Logger, basePath, rmHost string, store files.Storage) *Linguist {
	return &Linguist{l, service.NewBreakdown(l, basePath, rmHost, store)}
}

func (l *Linguist) Breakdown(ctx context.Context, rr *protos.BreakdownRequest) (*protos.BreakdownResponse, error) {
	l.log.Info("Handle request for Breakdown", "projectID", rr.GetProjectID(), "commit", rr.GetCommitHash())

	breakdownResult, err := l.breakdown.GetLanguages(rr.GetProjectID(), rr.GetCommitHash())
	if err != nil {
		return nil, err
	}

	return &protos.BreakdownResponse{Breakdown: breakdownResult}, nil
}
