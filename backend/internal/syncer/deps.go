package syncer

import (
	"context"

	"github.com/jmoiron/sqlx"
	content "github.com/mcherdakov/smart-drm/backend/internal/pkg/content/entity"
	proof "github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type drmService interface {
	SetProofs(ctx context.Context, proofs []proof.Proof) error
	SetClicks(ctx context.Context, clicks map[string]uint32) error
}

type proofsRepo interface {
	UnsyncedProofsAcquire(ctx context.Context, tx *sqlx.Tx) ([]proof.DBProof, error)
	UnsyncedProofsUpdate(ctx context.Context, tx *sqlx.Tx, proofs []proof.DBProof) error
	Transaction(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error
}

type statsRepo interface {
	UnsyncedClicksAcquire(ctx context.Context, tx *sqlx.Tx) ([]content.Click, error)
	UnsyncedClicksUpdate(ctx context.Context, tx *sqlx.Tx, proofs []content.Click) error
	Transaction(ctx context.Context, f func(ctx context.Context, tx *sqlx.Tx) error) error
}
