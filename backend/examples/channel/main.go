package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/mcherdakov/smart-drm/backend/internal/generated/channel"
)

const (
	contractAddressHex = "0x0165878A594ca255338adfa4d48449f69242Eb8F"
	// recieverAddressHex = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	recieverPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	chainURL        = "http://127.0.0.1:8545/"
	gasLimit uint64 = 300000
)

func pay(
	ctx context.Context,
	amount big.Int,
	ch *channel.Channel,
	client *ethclient.Client,
) error {
	privateKey, err := crypto.HexToECDSA(recieverPrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cast to ECDSA failed")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = &amount
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	tx, err := ch.Fund(auth)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, client, tx)

	return err
}

func run() error {
	ctx := context.Background()

	contractAddress := common.HexToAddress(contractAddressHex)
	client, err := ethclient.Dial(chainURL)
	if err != nil {
		return err
	}

	ch, err := channel.NewChannel(contractAddress, client)
	if err != nil {
		return err
	}

	balance, err := client.BalanceAt(ctx, contractAddress, nil)
	if err != nil {
		return err
	}

	fmt.Println((*balance).Int64())

	if err := pay(ctx, *big.NewInt(1000000), ch, client); err != nil {
		return err
	}

	balance, err = client.BalanceAt(ctx, contractAddress, nil)
	if err != nil {
		return err
	}

	fmt.Println((*balance).Int64())

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
