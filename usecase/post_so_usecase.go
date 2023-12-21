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
	var tempDistinct []string

	if marketplace == "Tokopedia" {
		pwMerchantPct = 0.045
		ongkirPct = 0.04
	}
	if marketplace == "Shopee" {
		pwMerchantPct = 0.04
		ongkirPct = 0.04
	}

	for _, saleOrder := range req {

		//JIKA SUDAH ADA DI DB, INVOICENYA DI SKIP
		orderbyId, _ := usecase.SbsRepository.GetSoById(c, saleOrder.InvoiceNo)
		if len(orderbyId) == 0 {
		} else {
			continue
		}

		//CARI KE DB BUAT DAPETIN HPP NYA
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
			IsPayment:        saleOrder.IsPayment,
		}

		errInsert := usecase.SbsRepository.PostSo(c, salesOrder)
		if errInsert != nil && errInsert.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errInsert != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errInsert.Error())
		}

		if saleOrder.Sku != "" {
			errUpdate := usecase.SbsRepository.DeductSbsProduct(c, saleOrder.Sku, saleOrder.Qty)
			if errUpdate != nil && errUpdate.Error() == config.ErrRecordNotFound.Error() {
				return response.BuildDataNotFoundResponse()
			} else if errUpdate != nil {
				return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errUpdate.Error())
			}
		}

		tempDistinct = append(tempDistinct, salesOrder.InvoiceNo)
	}

	errUpdateFlag := usecase.SbsRepository.UpdateSoFlag(c, tempDistinct)
	if errUpdateFlag != nil && errUpdateFlag.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errUpdateFlag != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errUpdateFlag.Error())
	}

	return response.BuildSuccessResponse(nil)

}
