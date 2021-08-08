package usecases

import (
	"context"
	"fmt"
	"rating-sekolah/domains"
	"time"
)

type schoolUsecase struct {
	schoolRepo    domains.SchoolRepository
	contextTimeout time.Duration
}

func NewSchoolUsecase(a domains.SchoolRepository, timeout time.Duration) domains.SchoolUseCase {
	return &schoolUsecase{
		schoolRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *schoolUsecase) Fetch(c context.Context, limit int64, offset int64) (res []domains.School, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.schoolRepo.Fetch(ctx, limit, offset)
	fmt.Println(res, "response")
	if err != nil {
		return nil, err
	}

	return
}

func (a *schoolUsecase) GetByID(c context.Context, id string) (res domains.School, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.schoolRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	return
}