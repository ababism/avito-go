package ports

import (
	"avito-go/services/banner/internal/domain"
	"context"
	"github.com/google/uuid"
)

type BannerRepository interface {
	Find(ctx context.Context, tag, feature int) (domain.Banner, error)

	Get(ctx context.Context, ID int) (domain.Banner, error)
	Create(ctx context.Context, course domain.Banner) (domain.Banner, error)
	GetList(ctx context.Context, filter domain.BannerFilter) ([]domain.Banner, error)
	Update(ctx context.Context, bannerID uuid.UUID, banner domain.Banner) (domain.Banner, error)
	Delete(ctx context.Context, bannerID uuid.UUID) error
	ProcessDeleteQueue(ctx context.Context) error
}

type BannerCache interface {
	UpdateMostViewed(ctx context.Context) error
	Get(ctx context.Context, tag, feature int) (domain.CachedBanner, error)
	Put(ctx context.Context, tag, feature int) (domain.CachedBanner, error)
}
