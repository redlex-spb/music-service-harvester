package usecase

import (
	"fmt"
	"github.com/redlex-spb/music-harvester/internal/gateway"
	"github.com/redlex-spb/music-harvester/pkg"
)

type ParseSongs struct{ prod gateway.Producer }

func NewParseSongs(p gateway.Producer) *ParseSongs { return &ParseSongs{p} }
func (uc *ParseSongs) Execute(r pkg.SongReq) error {
	if len(r.Songs) == 0 {
		return fmt.Errorf("empty")
	}
	return uc.prod.Send("puppeteer.parse", map[string]any{"type": "songs", "songs": r.Songs})
}
