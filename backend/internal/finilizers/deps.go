package finilizers

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type drmService interface {
	CallContractSplitBalance(ctx context.Context) (*types.Receipt, error)
	GetChannelFinishTimestamp(addr common.Address) (*time.Time, error)
	CallContractCloseChannel(
		ctx context.Context,
		hash []byte,
		channel common.Address,
	) (*types.Receipt, error)
}

type proofsRepo interface {
	Proofs(ctx context.Context) ([]entity.DBProof, error)
	DeleteProof(ctx context.Context, address string) error
}
