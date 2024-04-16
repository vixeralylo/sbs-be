package usecase

import (
	"context"
	"strconv"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) UpdateSbsProduct(c context.Context, sku string, quantity string) *response.ResponseContainer {

	// string to int
	qty, errConvert := strconv.Atoi(quantity)
	if errConvert != nil {
		// ... handle error
		panic(errConvert)
	}

	err := usecase.SbsRepository.UpdateSbsProduct(c, sku, qty)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(nil)

}
