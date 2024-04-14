package service

import (
	"avito-go/services/banner/internal/service/ports"
)

var _ ports.BannerService = &BannerService{}

type BannerService struct {
	cache ports.BannerCache

	repository ports.BannerRepository
}

func NewBannerService(bannerRepository ports.BannerRepository, bannerCache ports.BannerCache) ports.BannerService {
	return &BannerService{
		cache:      bannerCache,
		repository: bannerRepository,
	}
}

func (s BannerService) spanName(funcName string) string {
	return "banner/service." + funcName
}
