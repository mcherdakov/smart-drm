package content

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	contententity "github.com/mcherdakov/smart-drm/backend/internal/pkg/content/entity"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type proofsRepo interface {
	ProofByAddress(ctx context.Context, address string) (*entity.DBProof, error)
}

type DetailHandler struct {
	repo   contentRepo
	proofs proofsRepo
	drm    drmService
}

func NewDetailHandler(repo contentRepo, proofs proofsRepo, drm drmService) *DetailHandler {
	return &DetailHandler{
		repo:   repo,
		proofs: proofs,
		drm:    drm,
	}
}

type DetailItem struct {
	ID      int64  `json:"id"`
	Author  string `json:"author"`
	Header  string `json:"header"`
	Content string `json:"content"`
}

func (h *DetailHandler) Handle(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Error(fmt.Errorf("invalid id format: %w", err))
		return
	}

	address := c.Query("address")
	chanAddr, err := h.drm.UserAddrToChannelAddr(address)
	if err != nil {
		c.Error(err)
		return
	}

	proof, err := h.proofs.ProofByAddress(c, chanAddr)
	if err != nil {
		c.Error(err)
		return
	}

	if err := h.validateProof(proof); err != nil {
		c.Error(err)
		return
	}

	content, err := h.repo.ContentByID(c, id)
	if err != nil {
		c.Error(err)
		return
	}

	if content == nil {
		c.Error(fmt.Errorf("content with this id is not found"))
		return
	}

	click := contententity.Click{
		ContentID: content.ID,
		Date:      proof.Date,
		Address:   proof.Address,
	}

	if err := h.repo.CreateClick(c, click); err != nil {
		c.Error(err)
		return
	}

	response := DetailItem{
		ID:      content.ID,
		Author:  content.Author,
		Header:  content.Header,
		Content: content.Content,
	}

	c.JSON(http.StatusOK, response)
}

func (h *DetailHandler) validateProof(proof *entity.DBProof) error {
	if proof == nil {
		return fmt.Errorf("you haven't made payments through channel yet")
	}

	curDate := time.Now().Format(entity.DateFormat)
	if curDate != proof.Date {
		return fmt.Errorf("please pay for current date through your channel")
	}

	return nil
}
