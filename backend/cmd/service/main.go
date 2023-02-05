package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mcherdakov/smart-drm/backend/internal/services/ethereum"
)

func main() {
	address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	service, err := ethereum.NewService(os.Getenv("CHAIN_URL"), address)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	creators, err := service.GetCreators(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println(creators)
}
