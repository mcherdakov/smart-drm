package content

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/proofs/entity"
)

type ListHandler struct {
	repo contentRepo
	drm  drmService
}

func NewListHandler(repo contentRepo, drm drmService) *ListHandler {
	return &ListHandler{
		repo: repo,
		drm:  drm,
	}
}

type ListItem struct {
	ID          int64  `json:"id"`
	Author      string `json:"author"`
	Header      string `json:"header"`
	ClickExists bool   `json:"click_exists"`
}

func (h *ListHandler) Handle(c *gin.Context) {
	address := c.Query("address")
	chanAddr, err := h.drm.UserAddrToChannelAddr(address)
	if err != nil {
		c.Error(err)
		return
	}

	items, err := h.repo.Content(
		c,
		time.Now().Format(entity.DateFormat),
		chanAddr,
	)

	if err != nil {
		c.Error(err)
		return
	}

	response := make([]ListItem, 0, len(items))
	for _, item := range items {
		response = append(response, ListItem{
			ID:          item.ID,
			Author:      item.Author,
			Header:      item.Header,
			ClickExists: item.ClickExists,
		})
	}

	c.JSON(http.StatusOK, response)
}
