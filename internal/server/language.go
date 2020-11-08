package server

import (
	"context"

	"github.com/iantal/ld/internal/files"
	"github.com/iantal/ld/internal/service"
	"github.com/iantal/ld/internal/util"
	protos "github.com/iantal/ld/protos/ld"
	"github.com/sirupsen/logrus"
)

type Linguist struct {
	log       *util.StandardLogger
	breakdown *service.Breakdown
}

func NewLinguist(l *util.StandardLogger, basePath, rmHost string, store files.Storage) *Linguist {
	return &Linguist{l, service.NewBreakdown(l, basePath, rmHost, store)}
}

func (l *Linguist) Breakdown(ctx context.Context, rr *protos.BreakdownRequest) (*protos.BreakdownResponse, error) {
	l.log.WithFields(logrus.Fields{
		"projectID": rr.ProjectID,
		"commit": rr.CommitHash,
	}).Info("Handle request for Breakdown")

	breakdownResult, err := l.breakdown.GetLanguages(rr.GetProjectID(), rr.GetCommitHash())
	if err != nil {
		return nil, err
	}

	return &protos.BreakdownResponse{Breakdown: breakdownResult}, nil
}
