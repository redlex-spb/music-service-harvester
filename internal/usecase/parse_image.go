package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/internal/model"
	"time"
)

type ParseImage struct{ prod gateway.Producer }

func NewParseImage(p gateway.Producer) *ParseImage { return &ParseImage{p} }
func (uc *ParseImage) Execute(req model.ImageReq) error {
	if len(req.Bytes) == 0 {
		return fmt.Errorf("empty")
	}
	obj := "images/" + uuid.NewString() + "-" + req.Filename
	job := model.Job{ID: uuid.NewString(), Kind: model.JobImage, Payload: model.ImagePayload{Filename: req.Filename, ObjectPath: obj}, CreatedAt: time.Now(), UserID: req.UserID}
	return uc.prod.Send("harvest.jobs", job)
}
