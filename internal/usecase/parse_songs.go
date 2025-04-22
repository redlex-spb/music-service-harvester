package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/internal/model"
)

type ParseSongs struct{ prod gateway.Producer }

func NewParseSongs(p gateway.Producer) *ParseSongs { return &ParseSongs{p} }

func (uc *ParseSongs) Execute(req model.SongReq) error {
	job := model.Job{
		ID:        uuid.NewString(),
		Kind:      model.JobSongs,
		Payload:   model.SongPayload{Songs: req.Songs, Lang: req.Lang},
		CreatedAt: time.Now(),
		UserID:    req.UserID,
	}
	return uc.prod.Send("harvest.jobs", job)
}
