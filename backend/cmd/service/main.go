package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	drmHandler "github.com/mcherdakov/smart-drm/backend/internal/handlers/drm"
	"github.com/mcherdakov/smart-drm/backend/internal/handlers/pay"
	"github.com/mcherdakov/smart-drm/backend/internal/services/drm"
)

const contractAddress = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"

func run() error {
	drmService, err := drm.NewService(
		"http://127.0.0.1:8545/",
		common.HexToAddress(contractAddress[2:]),
	)
	if err != nil {
		return err
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	r.GET("/drm", drmHandler.NewHandler(contractAddress).Handle)
	r.POST("/pay", pay.NewHandler(drmService).Handle)

	return r.Run(":8000")
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
