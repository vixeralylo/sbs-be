package usecase

import (
	"context"
	"strconv"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) UpdateSbsProduct(c context.Context, sku string, quantity string, seqStr string, hpp float64, price float64) *response.ResponseContainer {

	// string to int
	qty, errConvert := strconv.Atoi(quantity)
	if errConvert != nil {
		// ... handle error
		panic(errConvert)
	}

	// seq opsional, default 0 kalau kosong
	seq := 0
	if seqStr != "" {
		seq, errConvert = strconv.Atoi(seqStr)
		if errConvert != nil {
			return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid seq")
		}
	}

	err := usecase.SbsRepository.UpdateSbsProduct(c, sku, qty, seq, hpp, price)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(nil)

}
