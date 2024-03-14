package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) UpdatePo(c context.Context, po_no string, status string) *response.ResponseContainer {

	err := usecase.SbsRepository.UpdatePo(c, po_no)

	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(nil)

}
