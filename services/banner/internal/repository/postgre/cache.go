package postgre

import (
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/service/ports"
	"context"
	"github.com/jmoiron/sqlx"
)

var _ ports.BannerCache = &bannerCache{}

type bannerCache struct {
	db *sqlx.DB
}

func NewBannerCache(db *sqlx.DB) ports.BannerCache {
	return &bannerCache{db: db}
}

func (b bannerCache) Get(ctx context.Context, tag, feature int) (domain.CachedBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerCache) Put(ctx context.Context, tag, feature int) (domain.CachedBanner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerCache) UpdateMostViewed(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
