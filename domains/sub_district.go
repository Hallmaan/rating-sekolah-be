package domains

import "context"

type SubDistrict struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SubDistrictUseCase interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]School, error)
	GetByID(ctx context.Context, id string) (School, error)
}

type SubDistrictRepository interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]School, error)
	GetByID(ctx context.Context, id string) (School, error)
}