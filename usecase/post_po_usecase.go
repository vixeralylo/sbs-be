package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) PostPo(c context.Context, req []dto.RequestPo) *response.ResponseContainer {

	for _, purchaseOrder := range req {

		purchasesOrder := entity.SbsPurchaseOrder{
			PoNumber:    purchaseOrder.PoNumber,
			PoDate:      purchaseOrder.PoDate,
			Sku:         purchaseOrder.Sku,
			ProductName: purchaseOrder.ProductName,
			Qty:         purchaseOrder.Qty,
			Price:       purchaseOrder.Price,
			Discount:    purchaseOrder.Discount,
			Ppn:         purchaseOrder.Ppn,
			TotalPrice:  purchaseOrder.TotalPrice,
		}

		errInsert := usecase.SbsRepository.PostPo(c, purchasesOrder)
		if errInsert != nil && errInsert.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errInsert != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errInsert.Error())
		}

		errUpdate := usecase.SbsRepository.AddSbsProduct(c, purchaseOrder.Sku, purchaseOrder.Qty)
		if errUpdate != nil && errUpdate.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errUpdate != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errUpdate.Error())
		}

	}

	return response.BuildSuccessResponse(nil)

}
