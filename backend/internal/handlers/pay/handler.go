package pay

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type drmService interface {
	ValidateProof(ctx context.Context, p entity.Proof) (string, error)
}

type proofsRepo interface {
	SetProof(ctx context.Context, p entity.Proof, address string) error
}

type Handler struct {
	drm  drmService
	repo proofsRepo
}

func NewHandler(drm drmService, repo proofsRepo) *Handler {
	return &Handler{
		drm:  drm,
		repo: repo,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	rawBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Error(err)
		return
	}

	var proof entity.Proof
	if err := json.Unmarshal(rawBody, &proof); err != nil {
		c.Error(err)
		return
	}

	address, err := h.drm.ValidateProof(c, proof)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.repo.SetProof(c, proof, address); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "ok",
	})
}
