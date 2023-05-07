package drm

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcherdakov/smart-drm/backend/internal/generated/smartdrm"
)

const (
	messagePrefix = "\x19Ethereum Signed Message:\n"
	gasLimit      = 10000000
)

type DRMService struct {
	client     *ethclient.Client
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey

	instance *smartdrm.Smartdrm
	address  common.Address
}

func NewService(url string, privateKey string) (*DRMService, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("cast to ECDSA failed")
	}

	return &DRMService{
		client:     client,
		privateKey: key,
		publicKey:  publicKeyECDSA,
	}, nil
}

func (s *DRMService) Address() common.Address {
	return s.address
}

func (s *DRMService) SetInstance(addr common.Address) error {
	instance, err := smartdrm.NewSmartdrm(addr, s.client)
	if err != nil {
		return err
	}

	s.instance = instance
	s.address = addr

	return nil
}

func (s *DRMService) DeployInstance(ctx context.Context) error {
	auth, err := s.makeAuth(ctx)
	if err != nil {
		return err
	}

	address, tx, instance, err := smartdrm.DeploySmartdrm(auth, s.client)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return err
	}

	s.instance = instance
	s.address = address

	return nil
}

func (s *DRMService) SetProof(ctx context.Context, p Proof) error {
	parsedProof, err := s.parseProof(p)
	if err != nil {
		return err
	}

	addr, err := s.userAddressFromProof(parsedProof, p.Hash)
	if err != nil {
		return err
	}

	channel, err := s.instance.GetUserChannel(nil, *addr)
	if err != nil {
		return err
	}

	chanProof, err := s.instance.GetChannelProof(nil, channel)
	if err != nil {
		return err
	}

	if p.Value-chanProof.Value.Int64() < SubscriptionPrice {
		return fmt.Errorf("invalid price")
	}

	proofs := []smartdrm.ChannelProof{
		{
			Proof:   *parsedProof,
			Channel: channel,
		},
	}

	return s.callContractSetProof(ctx, proofs)
}

func (s *DRMService) callContractSetProof(
	ctx context.Context,
	proofs []smartdrm.ChannelProof,
) error {
	auth, err := s.makeAuth(ctx)
	if err != nil {
		return err
	}

	tx, err := s.instance.SetChannelsProofs(auth, proofs)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(ctx, s.client, tx)
	return err
}

func (s *DRMService) makeAuth(ctx context.Context) (*bind.TransactOpts, error) {
	fromAddress := crypto.PubkeyToAddress(*s.publicKey)
	nonce, err := s.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(s.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func (s *DRMService) parseProof(p Proof) (*smartdrm.Proof, error) {
	rPart, err := hex.DecodeString(p.R[2:])
	if err != nil {
		return nil, err
	}

	sPart, err := hex.DecodeString(p.S[2:])
	if err != nil {
		return nil, err
	}

	return &smartdrm.Proof{
		V:     uint8(p.V),
		R:     *(*[32]byte)(rPart),
		S:     *(*[32]byte)(sPart),
		Value: big.NewInt(p.Value),
		Date:  p.Date,
	}, nil
}

func (s *DRMService) userAddressFromProof(p *smartdrm.Proof, h string) (*common.Address, error) {
	msgHash, err := hex.DecodeString(h[2:])
	if err != nil {
		return nil, err
	}

	prefixedHash := []byte(messagePrefix)
	prefixedHash = append(prefixedHash, []byte(fmt.Sprint(len(msgHash)))...)
	prefixedHash = append(prefixedHash, msgHash...)

	hash := crypto.Keccak256(prefixedHash)

	sig := []byte{}
	sig = append(sig, p.R[:]...)
	sig = append(sig, p.S[:]...)
	sig = append(sig, byte(p.V-27))

	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return nil, err
	}

	addr := crypto.PubkeyToAddress(*pubKey)

	return &addr, nil
}
