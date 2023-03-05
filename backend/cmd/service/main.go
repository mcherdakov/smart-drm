package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func test() error {
	privateKey := "df57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"

	bobKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)

	sig, err := crypto.Sign(hash.Bytes(), bobKey)
	if err != nil {
		return err
	}

	r := sig[:32]
	s := sig[:32:64]
	v := sig[64] + 27

	fmt.Println(r, s, v)

	return nil
}

func main() {
	if err := test(); err != nil {
		panic(err)
	}

	// address := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	// service, err := ethereum.NewService(os.Getenv("CHAIN_URL"), address)
	// if err != nil {
	// 	panic(err)
	// }

	// ctx := context.Background()
	// creators, err := service.GetCreators(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(creators)
}
