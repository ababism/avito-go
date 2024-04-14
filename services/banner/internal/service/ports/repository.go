package ports

import (
	"avito-go/services/banner/internal/domain"
	"context"
	"time"
)

type BannerRepository interface {
	Find(ctx context.Context, tag, feature int) (domain.Banner, error)

	Get(ctx context.Context, ID int) (domain.Banner, error)
	Create(ctx context.Context, course domain.Banner) (domain.Banner, error)
	GetList(ctx context.Context, filter domain.BannerFilter) ([]domain.Banner, error)
	Update(ctx context.Context, bannerID int, banner domain.Banner) (domain.Banner, error)
	Delete(ctx context.Context, bannerID int) error
	ProcessDeleteQueue(ctx context.Context) error
}

//type BannerCache interface {
//	UpdateMostViewed(ctx context.Context) error
//	Get(ctx context.Context, tag, feature int) (domain.CachedBanner, error)
//	Put(ctx context.Context, tag, feature int) (domain.CachedBanner, error)
//}

type BannerCache interface {
	Get(tag, feature int) (*domain.CachedBanner, bool)
	Delete(tag, feature int)
	Set(tag, feature int, banner domain.CachedBanner, duration time.Duration)
	Clean()
}
