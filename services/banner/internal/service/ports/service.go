package ports

import (
	"avito-go/services/banner/internal/domain"
	"context"
)

type BannerService interface {
	Find(ctx context.Context, actor domain.Actor, tag, feature int, useLastRevision bool) (domain.CachedBanner, error)

	Create(ctx context.Context, actor domain.Actor, course domain.Banner) (int, error)
	GetList(ctx context.Context, actor domain.Actor, filter domain.BannerFilter) ([]domain.Banner, error)
	Update(ctx context.Context, actor domain.Actor, bannerID int, banner domain.Banner) error
	Delete(ctx context.Context, actor domain.Actor, bannerID int) error

	//Get(ctx context.Context, actor domain.Actor, ID int) (domain.Banner, error)
}
