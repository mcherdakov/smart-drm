package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	r.GET("/bytecode", func(c *gin.Context) {
		data, err := os.ReadFile("bin/Channel.bin")
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"bytecode": string(data),
		})
	})

	r.Run(":8000")
}
