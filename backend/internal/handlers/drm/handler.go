package drm

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	contractAddress string
}

func NewHandler(contractAddress string) *Handler {
	return &Handler{
		contractAddress: contractAddress,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	data, err := os.ReadFile("abi/SmartDRM.abi")
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"abi":     string(data),
		"address": h.contractAddress,
	})
}
