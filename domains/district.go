package domains

import "context"

type District struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type DistrictUseCase interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]District, error)
	GetByID(ctx context.Context, id string) (District, error)
}

type DistrictRepository interface {
	Fetch(ctx context.Context, limit int64, offset int64) ([]District, error)
	GetByID(ctx context.Context, id string) (District, error)
}