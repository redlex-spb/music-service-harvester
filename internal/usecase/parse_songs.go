package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/internal/model"
	"time"
)

type ParseSongs struct{ prod gateway.Producer }

func NewParseSongs(p gateway.Producer) *ParseSongs { return &ParseSongs{p} }
func (uc *ParseSongs) Execute(req model.SongReq) error {
	if len(req.Songs) == 0 {
		return fmt.Errorf("empty")
	}
	job := model.Job{ID: uuid.NewString(), Kind: model.JobSongs, Payload: model.SongPayload{Songs: req.Songs, Lang: req.Lang}, CreatedAt: time.Now(), UserID: req.UserID}
	return uc.prod.Send("harvest.jobs", job)
}
