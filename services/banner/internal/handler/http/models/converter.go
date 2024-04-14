package models

import (
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/handler/http/banner/oapi"
	"encoding/json"
	"time"
)

const (
	UnknownVisibility = "unknown"
)

func ToRequestTime(t time.Time) *time.Time {
	var resT *time.Time
	if t == (time.Time{}) {
		resT = nil
	} else {
		resT = &t
	}
	return resT
}

func ToBannerResponse(banner domain.Banner) oapi.Banner {
	var content map[string]interface{}
	err := json.Unmarshal(banner.Content, &content)
	if err != nil {
		content = map[string]interface{}{}
	}
	// TODO add error
	return oapi.Banner{
		BannerId:  &banner.ID,
		Content:   &content,
		CreatedAt: ToRequestTime(banner.CreatedAt),
		FeatureId: &banner.Feature,
		IsActive:  &banner.IsActive,
		TagIds:    &banner.Tags,
		UpdatedAt: ToRequestTime(banner.UpdatedAt),
	}
}
func ToBannerListResponse(banner []domain.Banner) []oapi.Banner {
	var res []oapi.Banner
	for _, b := range banner {
		res = append(res, ToBannerResponse(b))
	}
	return res
}
