package usecase

import (
	"context"
	"sbs-be/model/dto"
	"sbs-be/model/response"
	"sbs-be/repository"
)

type SbsUsecase interface {
	GetSbsProduct(c context.Context) *response.ResponseContainer
	GetSo(c context.Context, filter dto.RequestSo) *response.ResponseContainer
	PostSo(c context.Context, marketplace string, req []dto.RequestContainer) *response.ResponseContainer
	PostPo(c context.Context, req []dto.RequestPo) *response.ResponseContainer
	UpdateSo(c context.Context, invoiceNo string, status string) *response.ResponseContainer
	GetSearchProduct(c context.Context, str string) *response.ResponseContainer
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
