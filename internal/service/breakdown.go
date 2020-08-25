package service

import (
	"github.com/hashicorp/go-hclog"
	"github.com/iantal/ld/protos/ld"
)

type Breakdown struct {
	log hclog.Logger
}

func NewBreakdown(l hclog.Logger) *Breakdown {
	return &Breakdown{l}
}

func (b *Breakdown) GetLanguages(projectID string) ([]*ld.Language, error) {
	return []*ld.Language{}, nil
}
