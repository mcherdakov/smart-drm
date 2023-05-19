package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
	"github.com/mcherdakov/smart-drm/backend/cmd/setup"
	"github.com/mcherdakov/smart-drm/backend/internal/generated/smartdrm"
	contractsRepo "github.com/mcherdakov/smart-drm/backend/internal/pkg/contracts/repository"
	"github.com/mcherdakov/smart-drm/backend/internal/services/drm"
)

func run() error {
	godotenv.Load()

	privateKey := os.Getenv("PRIVATE_KEY")
	chainAddress := os.Getenv("CHAIN_ADDRESS")

	ctx := context.Background()

	drmService, err := drm.NewService(
		chainAddress,
		privateKey,
	)
	if err != nil {
		return err
	}

	db, closeFunc, err := setup.DB()
	if err != nil {
		return err
	}

	contractRepo := contractsRepo.NewRepository(db)

	if err := setup.Contract(ctx, contractRepo, drmService); err != nil {
		return err
	}

	defer closeFunc()

	return testSet(ctx, drmService, 100)
}

func testSet(
	ctx context.Context,
	drmService *drm.DRMService,
	iterations int,
) error {
	maxCount := 20

	gasNewProof := make([]int64, maxCount)
	gasExistsProof := make([]int64, maxCount)

	gasNewClicks := make([]int64, maxCount)
	gasExistsClicks := make([]int64, maxCount)

	for r := 0; r < iterations; r++ {
		for i := 0; i < maxCount; i += 1 {
			addrs, err := genAddresses(i)
			if err != nil {
				return err
			}

			gasUsed, err := callSetProof(ctx, drmService, addrs, i, "2023-01-01")
			if err != nil {
				return err
			}
			gasNewProof[i] += gasUsed

			gasUsed, err = callSetProof(ctx, drmService, addrs, i, "2023-01-02")
			if err != nil {
				return err
			}
			gasExistsProof[i] += gasUsed

			gasUsed, err = callSetClicks(ctx, drmService, addrs, i)
			if err != nil {
				return err
			}
			gasNewClicks[i] += gasUsed

			gasUsed, err = callSetClicks(ctx, drmService, addrs, i)
			if err != nil {
				return err
			}
			gasExistsClicks[i] += gasUsed
		}
	}

	resNewProofs := make([]string, 0, maxCount)
	resExistsProofs := make([]string, 0, maxCount)
	resNewClicks := make([]string, 0, maxCount)
	resExistsClicks := make([]string, 0, maxCount)

	for i := 0; i < maxCount; i++ {
		resNewProofs = append(resNewProofs, fmt.Sprint(gasNewProof[i]/int64(iterations)))
		resExistsProofs = append(resExistsProofs, fmt.Sprint(gasExistsProof[i]/int64(iterations)))

		resNewClicks = append(resNewClicks, fmt.Sprint(gasNewClicks[i]/int64(iterations)))
		resExistsClicks = append(resExistsClicks, fmt.Sprint(gasExistsClicks[i]/int64(iterations)))
	}

	fmt.Println("new_proofs = [", strings.Join(resNewProofs, ", "), "]")
	fmt.Println("exists_proofs = [", strings.Join(resExistsProofs, ", "), "]")

	fmt.Println("new_clicks = [", strings.Join(resNewClicks, ", "), "]")
	fmt.Println("exists_clicks = [", strings.Join(resExistsClicks, ", "), "]")

	return nil
}

func callSetClicks(
	ctx context.Context,
	s *drm.DRMService,
	addrs []common.Address,
	count int,
) (int64, error) {
	clicks := make([]smartdrm.CreatorClicks, 0, count)

	for i := 0; i < count; i++ {
		clicks = append(clicks, smartdrm.CreatorClicks{
			Creator: addrs[i],
			Clicks:  rand.Uint32(),
		})
	}

	rec, err := s.CallContractSetCreatorClicks(ctx, clicks)
	if err != nil {
		return 0, err
	}

	return int64(rec.GasUsed), nil
}

func callSetProof(
	ctx context.Context,
	s *drm.DRMService,
	channels []common.Address,
	count int,
	date string,
) (int64, error) {
	proofs := make([]smartdrm.ChannelProof, 0, count)

	for i := 0; i < count; i++ {
		proofs = append(proofs, smartdrm.ChannelProof{
			Channel: channels[i],
			Proof: smartdrm.Proof{
				Value: big.NewInt(rand.Int63()),
				Date:  date,
			},
		})
	}

	rec, err := s.CallContractSetProof(ctx, proofs)
	if err != nil {
		return 0, err
	}

	return int64(rec.GasUsed), nil
}

func genAddresses(count int) ([]common.Address, error) {
	addrs := make([]common.Address, 0, count)
	for i := 0; i < count; i++ {
		priv, err := crypto.GenerateKey()
		if err != nil {
			return nil, err
		}

		pub := priv.Public().(*ecdsa.PublicKey)
		addr := crypto.PubkeyToAddress(*pub)
		addrs = append(addrs, addr)
	}

	return addrs, nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}

}
