package usecase

import (
	"context"
	"sbs-be/model/dto"
	"sbs-be/model/response"
	"sbs-be/repository"
)

type SbsUsecase interface {
	GetSbsProduct(c context.Context) *response.ResponseContainer
	UpdateSbsProduct(c context.Context, sku string, qty string) *response.ResponseContainer
	GetSo(c context.Context, filter dto.RequestSo) *response.ResponseContainer
	PostSo(c context.Context, marketplace string, req []dto.RequestContainer) *response.ResponseContainer
	PostSoManual(c context.Context, req dto.RequestSoManual) *response.ResponseContainer
	GetPo(c context.Context, filter dto.RequestPo) *response.ResponseContainer
	PostPo(c context.Context, req []dto.RequestPo) *response.ResponseContainer
	UpdateSo(c context.Context, invoiceNo string, status string) *response.ResponseContainer
	GetSearchProduct(c context.Context, str string) *response.ResponseContainer
	DeleteSo(c context.Context, invoiceNo string) *response.ResponseContainer
	UpdatePo(c context.Context, poNo string, status string) *response.ResponseContainer
	GetCost(c context.Context, filter dto.RequestCost) *response.ResponseContainer
	PostCost(c context.Context, req dto.RequestCost) *response.ResponseContainer
	GetSummary(c context.Context) *response.ResponseContainer
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
