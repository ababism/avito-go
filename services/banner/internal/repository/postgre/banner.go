package postgre

import (
	"avito-go/services/banner/internal/domain"
	"avito-go/services/banner/internal/service/ports"
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	spanBaseName = "banner/repository/postgres."
)

var _ ports.BannerRepository = &bannerRepository{}

type bannerRepository struct {
	db *sqlx.DB
}

func (b bannerRepository) ProcessDeleteQueue(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (b bannerRepository) Find(ctx context.Context, tag, feature int) (domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerRepository) Get(ctx context.Context, ID int) (domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerRepository) Create(ctx context.Context, course domain.Banner) (domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerRepository) GetList(ctx context.Context, filter domain.BannerFilter) ([]domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerRepository) Update(ctx context.Context, bannerID uuid.UUID, banner domain.Banner) (domain.Banner, error) {
	//TODO implement me
	panic("implement me")
}

func (b bannerRepository) Delete(ctx context.Context, bannerID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewBannerRepository(db *sqlx.DB) ports.BannerRepository {
	return &bannerRepository{db: db}
}
