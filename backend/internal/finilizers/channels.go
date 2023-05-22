package finilizers

import (
	"context"
	"encoding/hex"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

const (
	proofsSleepInterval = time.Second * 10
)

type ChannelsFinilizer struct {
	repo proofsRepo
	drm  drmService
}

func NewChannelsFinilizer(repo proofsRepo, drm drmService) *ChannelsFinilizer {
	return &ChannelsFinilizer{
		repo: repo,
		drm:  drm,
	}
}

func (s *ChannelsFinilizer) Run(ctx context.Context) error {
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

func (s *ChannelsFinilizer) process(ctx context.Context) error {
	proofs, err := s.repo.Proofs(ctx)
	if err != nil {
		return err
	}

	for _, proof := range proofs {
		if err := s.handleProof(ctx, proof); err != nil {
			return err
		}
	}

	return nil
}

func (s *ChannelsFinilizer) handleProof(ctx context.Context, proof entity.DBProof) error {
	addr := common.HexToAddress(proof.Address)
	finishTime, err := s.drm.GetChannelFinishTimestamp(addr)
	if err != nil {
		return err
	}

	safeTime := time.Now().Add(time.Hour)
	if finishTime.Before(safeTime) {
		hash, err := hex.DecodeString(proof.Hash[2:])
		if err != nil {
			return err
		}

		_, err = s.drm.CallContractCloseChannel(ctx, hash, addr)
		if err != nil {
			return err
		}

		if err := s.repo.DeleteProof(ctx, proof.Address); err != nil {
			return err
		}
	}

	return nil
}
