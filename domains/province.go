package domains

import "context"

type Province struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProvinceUseCase interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]Province, error)
	GetByID(ctx context.Context, id string) (Province, error)
}

type ProvinceRepository interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]Province, error)
	GetByID(ctx context.Context, id string) (Province, error)
}