package stats

import (
	"context"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/mcherdakov/smart-drm/backend/internal/pkg/content/entity"
)

type drmService interface {
}

type contentRepo interface {
	AuthorStats(ctx context.Context) ([]entity.AuthorClickAggregate, error)
	ContentStats(ctx context.Context) ([]entity.ContentClickAggregate, error)
}

type Handler struct {
	repo contentRepo
	drm  drmService
}

func NewHandler(repo contentRepo, drm drmService) *Handler {
	return &Handler{
		repo: repo,
		drm:  drm,
	}
}

type AuthorStats struct {
	Author           string `json:"author"`
	ClicksCommited   int64  `json:"clicks_commited"`
	ClicksUncommited int64  `json:"clicks_uncommited"`
}

type ContentStats struct {
	ID               int64  `json:"id"`
	Author           string `json:"author"`
	Header           string `json:"header"`
	ClicksCommited   int64  `json:"clicks_commited"`
	ClicksUncommited int64  `json:"clicks_uncommited"`
}

type Response struct {
	Author  []AuthorStats  `json:"author"`
	Content []ContentStats `json:"content"`
}

func (h *Handler) Handle(c *gin.Context) {
	authorStats, err := h.repo.AuthorStats(c)
	if err != nil {
		c.Error(err)
		return
	}

	contentStats, err := h.repo.ContentStats(c)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Author:  h.authorToResult(authorStats),
		Content: h.contentToResult(contentStats),
	})
}

func (h *Handler) authorToResult(author []entity.AuthorClickAggregate) []AuthorStats {
	resMap := map[string]AuthorStats{}
	for _, a := range author {
		stats := resMap[a.Author]

		stats.Author = a.Author

		if a.Commited {
			stats.ClicksCommited = a.Count
		} else {
			stats.ClicksUncommited = a.Count
		}

		resMap[a.Author] = stats
	}

	res := make([]AuthorStats, 0, len(resMap))
	for _, r := range resMap {
		res = append(res, r)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].ClicksUncommited < res[j].ClicksUncommited
	})

	return res
}

func (h *Handler) contentToResult(content []entity.ContentClickAggregate) []ContentStats {
	res := make([]ContentStats, 0, len(content))

	for _, c := range content {
		res = append(res, ContentStats{
			ID:               c.ID,
			Author:           c.Author,
			Header:           c.Header,
			ClicksCommited:   c.ClicksCommited,
			ClicksUncommited: c.ClicksUncommited,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].ClicksUncommited < res[j].ClicksUncommited
	})

	return res
}
