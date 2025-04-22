package usecase

import (
	"github.com/redlex-spb/music-harvester/internal/model"
	"testing"
)

func TestParseSources(t *testing.T) {
	mp := &mockProducer{}
	uc := NewParseSources(mp)
	if err := uc.Execute(model.SourceReq{URLs: []string{"https://x"}}); err != nil {
		t.Fatalf("unexpected %v", err)
	}
	if mp.topic != "harvest.jobs" {
		t.Errorf("topic %s", mp.topic)
	}
}
