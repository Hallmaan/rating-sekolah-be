package usecases

import (
	"context"
	"fmt"
	"rating-sekolah/domains"
	"time"
)

type provinceUsecase struct {
	provinceRepo    domains.ProvinceRepository
	contextTimeout time.Duration
}

func NewProvinceUsecase(a domains.ProvinceRepository, timeout time.Duration) domains.ProvinceUseCase {
	return &provinceUsecase{
		provinceRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *provinceUsecase) Fetch(c context.Context, limit int64, offset int64) (res []domains.Province, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.provinceRepo.Fetch(ctx, limit, offset)
	fmt.Println(res, "response")
	if err != nil {
		return nil, err
	}

	return
}

func (a *provinceUsecase) GetByID(c context.Context, id string) (res domains.Province, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.provinceRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}
