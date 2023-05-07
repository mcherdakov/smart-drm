package pay

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mcherdakov/smart-drm/backend/internal/services/drm"
)

type drmService interface {
	SetProof(ctx context.Context, p drm.Proof) error
}

type Handler struct {
	drm drmService
}

func NewHandler(drm drmService) *Handler {
	return &Handler{
		drm: drm,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	rawBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
		return
	}

	var proof drm.Proof
	if err := json.Unmarshal(rawBody, &proof); err != nil {
		c.Error(err)
		return
	}

	if err := h.drm.SetProof(c, proof); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
