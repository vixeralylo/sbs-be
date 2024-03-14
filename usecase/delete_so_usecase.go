package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) DeleteSo(c context.Context, invoice_no string) *response.ResponseContainer {

	filter := dto.RequestSo{
		SoNumber: invoice_no,
	}
	resultOrder, errGetSo := usecase.SbsRepository.GetSo(c, filter)
	if errGetSo != nil && errGetSo.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errGetSo != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errGetSo.Error())
	}

	for _, salesOrder := range resultOrder {
		sku := salesOrder.Sku
		qty := salesOrder.Qty

		errAddStock := usecase.SbsRepository.AddSbsProduct(c, sku, qty)
		if errAddStock != nil && errAddStock.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errAddStock != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errAddStock.Error())
		}
	}

	err := usecase.SbsRepository.UpdateSoCancel(c, invoice_no)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(nil)

}
