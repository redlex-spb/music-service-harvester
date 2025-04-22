package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/internal/model"
)

type ParseSources struct{ prod gateway.Producer }

func NewParseSources(p gateway.Producer) *ParseSources { return &ParseSources{p} }

func (uc *ParseSources) Execute(req model.SourceReq) error {
	job := model.Job{
		ID:        uuid.NewString(),
		Kind:      model.JobSources,
		Payload:   model.SourcePayload{URLs: req.URLs, Lang: req.Lang},
		CreatedAt: time.Now(),
		UserID:    req.UserID,
	}
	return uc.prod.Send("harvest.jobs", job)
}
