package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) PostSo(c context.Context, marketplace string, req []dto.RequestContainer) *response.ResponseContainer {

	var pwMerchantPct float32
	var ongkirPct float32

	if marketplace == "Tokopedia" {
		pwMerchantPct = 0.045
		ongkirPct = 0.04
	}
	if marketplace == "Shopee" {
		pwMerchantPct = 0.04
		ongkirPct = 0.04
	}

	for _, saleOrder := range req {

		productById, err := usecase.SbsRepository.GetSbsProductById(c, saleOrder.Sku)
		if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if err != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
		}

		var hpp int
		var productName string
		if saleOrder.Sku == "" {
			hpp = 0
			productName = "Plastik Bubble/ Bubble Wrap Packing tambahan"
		} else {
			hpp = productById[0].Hpp
			productName = productById[0].ProductName
		}

		totalPrice := saleOrder.Qty * saleOrder.Price
		pwMerchantFee := pwMerchantPct * float32(totalPrice)
		ongkirFee := ongkirPct * float32(totalPrice)
		grossMargin := totalPrice - (hpp * saleOrder.Qty)
		cleanMargin := float32(grossMargin) - pwMerchantFee - ongkirFee

		salesOrder := entity.SbsSalesOrder{
			MarketPlaceId:    marketplace,
			OrderDate:        saleOrder.OrderDate,
			InvoiceNo:        saleOrder.InvoiceNo,
			Sku:              saleOrder.Sku,
			ProductName:      productName,
			Qty:              saleOrder.Qty,
			SalesPrice:       saleOrder.Price,
			TotalPrice:       totalPrice,
			Hpp:              hpp * saleOrder.Qty,
			GrossMargin:      grossMargin,
			PowerMerchantFee: pwMerchantFee,
			OngkirFee:        ongkirFee,
			CleanMargin:      cleanMargin,
		}

		errInsert := usecase.SbsRepository.PostSo(c, salesOrder)
		if errInsert != nil && errInsert.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errInsert != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errInsert.Error())
		}

		if saleOrder.Sku != "" {
			errUpdate := usecase.SbsRepository.UpdateSbsProduct(c, saleOrder.Sku, saleOrder.Qty)
			if errUpdate != nil && errUpdate.Error() == config.ErrRecordNotFound.Error() {
				return response.BuildDataNotFoundResponse()
			} else if errUpdate != nil {
				return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errUpdate.Error())
			}
		}

	}

	return response.BuildSuccessResponse(nil)

}
