package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) PostCost(c context.Context, req dto.RequestCost) *response.ResponseContainer {

	filter := entity.SbsCost{
		Date:          req.Date,
		CostType:      req.CostType,
		CostName:      req.CostName,
		Qty:           req.Qty,
		Price:         req.Price,
		AddedPrice:    req.AddedPrice,
		TotalPrice:    req.TotalPrice,
		MarketplaceId: req.MarketplaceId,
		InvoiceNo:     req.InvoiceNo,
	}
	err := usecase.SbsRepository.PostCost(c, filter)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(nil)

}
