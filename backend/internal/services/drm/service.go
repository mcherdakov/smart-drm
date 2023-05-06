package drm

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mcherdakov/smart-drm/backend/internal/generated/smartdrm"
)

const messagePrefix = "\x19Ethereum Signed Message:\n"

type DRMService struct {
	client   *ethclient.Client
	instance *smartdrm.Smartdrm
}

func NewService(url string, contractAddr common.Address) (*DRMService, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	instance, err := smartdrm.NewSmartdrm(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return &DRMService{
		client:   client,
		instance: instance,
	}, nil
}

func (s *DRMService) SetProof(p Proof) error {
	msgHash, err := hex.DecodeString(p.Hash[2:])
	if err != nil {
		return err
	}

	prefixedHash := []byte(messagePrefix)
	prefixedHash = append(prefixedHash, []byte(fmt.Sprint(len(msgHash)))...)
	prefixedHash = append(prefixedHash, msgHash...)

	hash := crypto.Keccak256(prefixedHash)

	rPart, err := hex.DecodeString(p.R[2:])
	if err != nil {
		return err
	}

	sPart, err := hex.DecodeString(p.S[2:])
	if err != nil {
		return err
	}

	sig := []byte{}
	sig = append(sig, rPart...)
	sig = append(sig, sPart...)
	sig = append(sig, byte(p.V-27))

	fmt.Println(msgHash)

	pubKey, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return err
	}

	addr := crypto.PubkeyToAddress(*pubKey)
	fmt.Println(addr)

	return nil
}
