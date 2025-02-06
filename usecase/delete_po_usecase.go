package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) DeletePo(c context.Context, poNo string, sku string) *response.ResponseContainer {

	filter := dto.RequestPo{
		PoNumber: poNo,
		Sku:      sku,
	}
	resultOrder, errGetPo := usecase.SbsRepository.GetPo(c, filter)
	if errGetPo != nil && errGetPo.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errGetPo != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errGetPo.Error())
	}

	errAddStock := usecase.SbsRepository.DeductSbsProduct(c, sku, resultOrder[0].Qty)
	if errAddStock != nil && errAddStock.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errAddStock != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errAddStock.Error())
	}

	errRemovePo := usecase.SbsRepository.DeletePo(c, poNo, sku)
	if errRemovePo != nil && errRemovePo.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errRemovePo != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errRemovePo.Error())
	}

	return response.BuildSuccessResponse(nil)

}
