package domain

import (
	"avito-go/pkg/xstringset"
	"encoding/json"
	"time"
)

const (
	ServiceName = "avito-go/banner"
)

// Actor represents a user of system with data taken from request
type Actor struct {
	//ID    uuid.UUID
	roles xstringset.Set
}

type Banner struct {
	//UUID      uuid.UUID
	ID int

	json.RawMessage
	//Content *map[string]interface{}

	Content  json.RawMessage
	IsActive bool

	Feature int
	Tags    []int

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CachedBanner struct {
	ID int

	//Content *map[string]interface{}
	Content json.RawMessage
	
	IsActive bool

	Expiration int64
	//CreatedAt time.Time
}

type BannerFilter struct {
	Feature *int
	TagID   *int

	Limit  *int
	Offset *int
}
