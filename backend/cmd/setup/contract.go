package setup

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/contracts/entity"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/contracts/repository"
	"github.com/mcherdakov/smart-drm/backend/internal/services/drm"
)

func Contract(
	ctx context.Context,
	contractRepo *repository.Repository,
	service *drm.DRMService,
) error {
	dbContract, err := contractRepo.ContractByName(ctx, entity.ContractNameSmartDRM)
	if err != nil {
		return err
	}

	if dbContract != nil {
		service.SetInstance(common.HexToAddress(dbContract.Address))

		return nil
	}

	if err := service.DeployInstance(ctx); err != nil {
		return err
	}

	return contractRepo.ContractInsert(
		ctx,
		entity.Contract{
			Name:    entity.ContractNameSmartDRM,
			Address: service.Address().String(),
		},
	)
}
