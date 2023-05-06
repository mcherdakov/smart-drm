package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const contractAddress = "0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0"

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	r.GET("/drm", func(c *gin.Context) {
		data, err := os.ReadFile("abi/SmartDRM.abi")
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"abi":     string(data),
			"address": contractAddress,
		})
	})

	r.Run(":8000")
}
