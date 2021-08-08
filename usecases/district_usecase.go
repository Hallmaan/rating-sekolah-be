package usecases

import (
	"context"
	"fmt"
	"rating-sekolah/domains"
	"time"
)

type districtUsecase struct {
	districtRepo    domains.DistrictRepository
	contextTimeout time.Duration
}

func NewDistrictUsecase(a domains.DistrictRepository, timeout time.Duration) domains.DistrictUseCase {
	return &districtUsecase{
		districtRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *districtUsecase) Fetch(c context.Context, limit int64, offset int64) (res []domains.District, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.districtRepo.Fetch(ctx, limit, offset)
	fmt.Println(res, "response")
	if err != nil {
		return nil, err
	}

	return
}

func (a *districtUsecase) GetByID(c context.Context, id string) (res domains.District, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.districtRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}
