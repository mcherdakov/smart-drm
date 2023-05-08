package syncer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

const (
	proofsSleepInterval    = time.Second * 10
	paymentsCountThreshold = 3
)

type drmService interface {
	SetProofs(ctx context.Context, proofs []entity.Proof) error
}

type proofsRepo interface {
	UnsyncedProofsAcquire(ctx context.Context, tx *sqlx.Tx) ([]entity.DBProof, error)
	UnsyncedProofsUpdate(ctx context.Context, tx *sqlx.Tx, proofs []entity.DBProof) error
	Transaction(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error
}

type ProofsSyncer struct {
	repo proofsRepo
	drm  drmService
}

func NewProofsSyncer(repo proofsRepo, drm drmService) *ProofsSyncer {
	return &ProofsSyncer{
		repo: repo,
		drm:  drm,
	}
}

func (s *ProofsSyncer) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := s.process(ctx); err != nil {
			return err
		}

		time.Sleep(proofsSleepInterval)
	}
}

func (s *ProofsSyncer) process(ctx context.Context) error {
	return s.repo.Transaction(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		proofs, err := s.repo.UnsyncedProofsAcquire(ctx, tx)
		if err != nil {
			return err
		}

		paymentsCount := 0
		for _, proof := range proofs {
			paymentsCount += int(calcDiff(proof))
		}

		log.Println(fmt.Sprintf("found %d payments", paymentsCount))

		if paymentsCount < paymentsCountThreshold {
			return nil
		}

		contractProofs := make([]entity.Proof, 0, len(proofs))
		for _, p := range proofs {
			contractProofs = append(contractProofs, entity.Proof{
				V:     p.V,
				R:     p.R,
				S:     p.S,
				Hash:  p.Hash,
				Date:  p.Date,
				Value: p.Value,
			})
		}

		if err := s.repo.UnsyncedProofsUpdate(ctx, tx, proofs); err != nil {
			return err
		}

		return s.drm.SetProofs(ctx, contractProofs)
	})
}

func calcDiff(p entity.DBProof) int64 {
	diffValues := p.Value - p.LastCommitedValue.ValueOrZero()

	return diffValues / entity.SubscriptionPrice
}
