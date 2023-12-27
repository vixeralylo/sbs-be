package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) GetSo(c context.Context, filter dto.RequestSo) *response.ResponseContainer {

	resultOrder, err := usecase.SbsRepository.GetSo(c, filter)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	var sumPrice int
	var sumTotalPrice float64
	var sumHpp float64
	var sumMargin float64
	var sumAdmin float64
	var sumOngkir float64
	var sumCleanMargin float64

	for _, order := range resultOrder {
		sumPrice = sumPrice + order.SalesPrice
		sumTotalPrice = sumTotalPrice + float64(order.TotalPrice)
		sumHpp = sumHpp + float64(order.Hpp)
		sumMargin = sumMargin + float64(order.GrossMargin)
		sumAdmin = sumAdmin + float64(order.PowerMerchantFee)
		sumOngkir = sumOngkir + float64(order.OngkirFee)
		sumCleanMargin = sumCleanMargin + float64(order.CleanMargin)
	}

	result := entity.SbsSalesOrderResponse{
		SumPrice:       sumPrice,
		SumTotalPrice:  sumTotalPrice,
		SumHpp:         sumHpp,
		SumMargin:      sumMargin,
		SumAdmin:       sumAdmin,
		SumOngkir:      sumOngkir,
		SumCleanMargin: sumCleanMargin,
		SbsOrderList:   resultOrder,
	}

	return response.BuildSuccessResponse(result)

}
