package usecase

import (
	"fmt"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/pkg"
)

type ParseImage struct{ prod gateway.Producer }

func NewParseImage(p gateway.Producer) *ParseImage { return &ParseImage{p} }
func (uc *ParseImage) Execute(r pkg.ImageReq) error {
	if len(r.Bytes) == 0 {
		return fmt.Errorf("empty")
	}
	return uc.prod.Send("puppeteer.parse", map[string]any{"type": "image", "file": r.Filename})
}
