package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	drmHandler "github.com/mcherdakov/smart-drm/backend/internal/handlers/drm"
	"github.com/mcherdakov/smart-drm/backend/internal/handlers/pay"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/contracts/repostory"
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

	db, closeFunc, err := setupDB()
	if err != nil {
		return err
	}
	defer closeFunc()

	contractRepo := repostory.NewRepository(db)

	if err := setupContract(ctx, contractRepo, drmService); err != nil {
		return err
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	r.GET("/drm", drmHandler.NewHandler(drmService.Address()).Handle)
	r.POST("/pay", pay.NewHandler(drmService).Handle)

	return r.Run(":8000")
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
