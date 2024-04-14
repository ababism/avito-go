package models

import (
	"avito-go/services/banner/internal/domain"
)

func (b *Banner) ToDomain() domain.Banner {
	return domain.Banner{
		ID:        b.ID,
		Content:   b.Content,
		IsActive:  b.IsActive,
		Feature:   b.Feature,
		Tags:      b.Tags,
		UpdatedAt: b.UpdatedAt,
		CreatedAt: b.CreatedAt,
	}
}

func ToBannerModel(b domain.Banner) Banner {
	return Banner{
		ID:        b.ID,
		Content:   b.Content,
		IsActive:  b.IsActive,
		Feature:   b.Feature,
		Tags:      b.Tags,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
