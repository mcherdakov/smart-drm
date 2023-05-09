package content

import (
	"context"

	"github.com/mcherdakov/smart-drm/backend/internal/pkg/content/entity"
)

type contentRepo interface {
	ContentByID(ctx context.Context, id int64) (*entity.Content, error)
	Content(ctx context.Context, date string, address string) ([]entity.Content, error)
	CreateClick(ctx context.Context, click entity.Click) error
}

type drmService interface {
	UserAddrToChannelAddr(userAddr string) (string, error)
}
