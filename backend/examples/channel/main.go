package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcherdakov/smart-drm/backend/internal/generated/channel"
)

const (
	contractAddressHex = "Cf7Ed3AccA5a467e9e704C703E8D87F634fB0Fc9"
	recieverPrivateKey = "df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"

	chainURL        = "http://127.0.0.1:8545/"
	gasLimit uint64 = 300000
)

func create(
	ctx context.Context,
	contractAddress common.Address,
	client *ethclient.Client,
) (*channel.Channel, error) {
	return channel.NewChannel(contractAddress, client)
}

func run() error {
	ctx := context.Background()

	contractAddress := common.HexToAddress(contractAddressHex)

	client, err := ethclient.Dial(chainURL)
	if err != nil {
		return err
	}

	_, err = create(ctx, contractAddress, client)
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(ctx, contractAddress, nil)
	if err != nil {
		return err
	}

	fmt.Println("starting balance is", (*balance).Int64())

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
