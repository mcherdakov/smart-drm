package syncer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	statsSleepInterval = time.Second * 10
)

type StatsSyncer struct {
	repo statsRepo
	drm  drmService
}

func NewStatsSyncer(repo statsRepo, drm drmService) *StatsSyncer {
	return &StatsSyncer{
		repo: repo,
		drm:  drm,
	}
}

func (s *StatsSyncer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := s.process(ctx); err != nil {
			return err
		}

		time.Sleep(statsSleepInterval)
	}
}

func (s *StatsSyncer) process(ctx context.Context) error {
	return s.repo.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		clicks, err := s.repo.UnsyncedClicksAcquire(ctx, tx)
		if err != nil {
			return err
		}

		clicksAgg := map[string]uint32{}
		totalClicks := 0

		for i := range clicks {
			clicks[i].Commited = true

			clicksAgg[clicks[i].Address]++
			totalClicks++
		}

		log.Println(fmt.Sprintf("found %d clicks", totalClicks))

		if totalClicks == 0 {
			return nil
		}

		if err := s.repo.UnsyncedClicksUpdate(ctx, tx, clicks); err != nil {
			return err
		}

		return s.drm.SetClicks(ctx, clicksAgg)
	})
}
