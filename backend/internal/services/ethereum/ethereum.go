package ethereum

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcherdakov/smart-drm/backend/internal/generated/smartdrm"
)

type Service interface{}

type GoEthService struct {
	client   *ethclient.Client
	instance *smartdrm.Smartdrm
}

func NewService(url string, contractAddr common.Address) (*GoEthService, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	instance, err := smartdrm.NewSmartdrm(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &GoEthService{
		client:   client,
		instance: instance,
	}, nil
}

func (s *GoEthService) GetCreators(ctx context.Context) ([]common.Address, error) {
	var index int64
	addrs := []common.Address{}

	for {
		addr, err := s.instance.GetCreator(nil, big.NewInt(index))

		if err != nil {
			if strings.Contains(err.Error(), "SmartDRMOutOfBounds") {
				break
			}

			return nil, err
		}

		addrs = append(addrs, addr)
		index++
	}

	return addrs, nil
}
