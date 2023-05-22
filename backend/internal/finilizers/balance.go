package finilizers

import (
	"context"
	"time"
)

const (
	balanceSleepInterval = time.Hour * 24
)

type BalanceFinilizer struct {
	drm drmService
}

func NewBalanceFinilizer(drm drmService) *BalanceFinilizer {
	return &BalanceFinilizer{
		drm: drm,
	}
}

func (s *BalanceFinilizer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := s.process(ctx); err != nil {
			return err
		}

		time.Sleep(balanceSleepInterval)
	}
}

func (s *BalanceFinilizer) process(ctx context.Context) error {
	_, err := s.drm.CallContractSplitBalance(ctx)

	return err
}
