package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"

	drmHandler "github.com/mcherdakov/smart-drm/backend/internal/handlers/drm"
	"github.com/mcherdakov/smart-drm/backend/internal/handlers/pay"
	"github.com/mcherdakov/smart-drm/backend/internal/handlers/proof"
	contractsRepo "github.com/mcherdakov/smart-drm/backend/internal/pkg/contracts/repository"
	proofsrepo "github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/repository"
	"github.com/mcherdakov/smart-drm/backend/internal/services/drm"
	"github.com/mcherdakov/smart-drm/backend/internal/syncer"
)

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) != 0 {
		c.JSON(http.StatusInternalServerError, c.Errors)
	}
}

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

	db, closeFunc, err := setupDB()
	if err != nil {
		return err
	}

	defer closeFunc()

	contractRepo := contractsRepo.NewRepository(db)
	proofsRepo := proofsrepo.NewRepository(db)

	if err := setupContract(ctx, contractRepo, drmService); err != nil {
		return err
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))
	r.Use(errorHandler)

	r.GET("/drm", drmHandler.NewHandler(drmService.Address()).Handle)
	r.GET("/proof", proof.NewHandler(proofsRepo).Handle)
	r.POST("/pay", pay.NewHandler(drmService, proofsRepo).Handle)

	errg := errgroup.Group{}

	errg.Go(func() error {
		return r.Run(":8000")
	})

	errg.Go(func() error {
		err := syncer.NewProofsSyncer(proofsRepo, drmService).Run(ctx)
		if err != nil {
			log.Println(err)
		}
		return err
	})

	return errg.Wait()
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
