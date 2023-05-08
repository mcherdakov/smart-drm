package proof

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type proofsRepo interface {
	ProofByAddress(ctx context.Context, address string) (*entity.DBProof, error)
}

type Handler struct {
	repo proofsRepo
}

func NewHandler(repo proofsRepo) *Handler {
	return &Handler{
		repo: repo,
	}
}

type Request struct {
	Address string `json:"address"`
}

func (h *Handler) Handle(c *gin.Context) {
	address := strings.ToLower(c.Query("address"))

	proof, err := h.repo.ProofByAddress(c, address)
	if err != nil {
		c.Error(err)
		return
	}

	if proof == nil {
		proof = &entity.DBProof{}
	}

	c.JSON(http.StatusOK, proof)
}
