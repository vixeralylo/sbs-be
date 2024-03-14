package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) GetCost(c context.Context, filter dto.RequestCost) *response.ResponseContainer {

	result, err := usecase.SbsRepository.GetCost(c, filter)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(result)

}
