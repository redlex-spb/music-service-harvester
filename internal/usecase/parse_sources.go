package usecase

import (
	"fmt"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/pkg"
)

type ParseSources struct{ prod gateway.Producer }

func NewParseSources(p gateway.Producer) *ParseSources { return &ParseSources{p} }
func (uc *ParseSources) Execute(r pkg.SourceReq) error {
	if len(r.URLs) == 0 {
		return fmt.Errorf("empty")
	}
	return uc.prod.Send("puppeteer.parse", map[string]any{"type": "sources", "urls": r.URLs})
}
