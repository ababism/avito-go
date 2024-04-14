package oapi

import (
	"avito-go/services/banner/internal/domain"
	"encoding/json"
	"errors"
)

func (b PostBannerJSONBody) ToDomain() (domain.Banner, error) {
	if b.Content == nil {
		return domain.Banner{}, errors.New("content is required")
	}
	if b.IsActive == nil {
		return domain.Banner{}, errors.New("isActive is required")
	}
	if b.FeatureId == nil {
		return domain.Banner{}, errors.New("feature is required")
	}
	if b.TagIds == nil {
		return domain.Banner{}, errors.New("tags is required")
	}
	content, err := json.Marshal(b.Content)
	if err != nil {
		return domain.Banner{}, err
	}

	if len(*b.TagIds) == 0 {
		return domain.Banner{}, errors.New("at least one tag is required")
	}
	return domain.Banner{
		ID:       0,
		Content:  content,
		IsActive: *b.IsActive,
		Feature:  *b.FeatureId,
		Tags:     *b.TagIds,
	}, nil
}

func (f GetBannerParams) ToDomain() domain.BannerFilter {
	return domain.BannerFilter{
		Feature: f.FeatureId,
		TagID:   f.TagId,
		Limit:   f.Limit,
		Offset:  f.Offset,
	}
}

//func (a Actor) ToDomainWithRoles(roles []string) domain.Actor {
//	roles = append(roles, a.Roles)
//	return domain.NewActor(a.ID, roles)
//}
//func ToTimeDomain(t *time.Time) time.Time {
//	if t == nil {
//		return time.Time{}
//	}
//	return *t
//}
//func (a Actor) ToDomain() domain.Actor {
//	s := []string{a.Roles}
//	return domain.NewActor(a.ID, s)
//}
