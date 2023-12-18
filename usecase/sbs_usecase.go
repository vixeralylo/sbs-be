package usecase

import (
	"context"
	"sbs-be/model/dto"
	"sbs-be/model/response"
	"sbs-be/repository"
)

type SbsUsecase interface {
	GetSbsProduct(c context.Context) *response.ResponseContainer
	PostSo(c context.Context, marketplace string, req []dto.RequestContainer) *response.ResponseContainer
	Ping() string
}

type sbsUsecase struct {
	SbsRepository repository.SbsRepository
}

func GetSbsUsecase(sbsRepository repository.SbsRepository) SbsUsecase {
	return &sbsUsecase{
		SbsRepository: sbsRepository,
	}
}

func (usecase *sbsUsecase) Ping() string {
	return usecase.SbsRepository.Ping()
}
