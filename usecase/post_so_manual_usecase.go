package usecase

import (
	"context"
	"strconv"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) PostSoManual(c context.Context, req dto.RequestSoManual) *response.ResponseContainer {

	//CARI KE DB BUAT DAPETIN HPP NYA
	productById, err := usecase.SbsRepository.GetSbsProductById(c, req.Sku)
	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	var hpp int
	var productName string
	if req.Sku == "" {
		hpp = 0
		productName = "Plastik Bubble/ Bubble Wrap Packing tambahan"
	} else {
		hpp = productById[0].Hpp
		productName = productById[0].ProductName
	}

	qty, errConvert := strconv.Atoi(req.Qty)
	if errConvert != nil && errConvert.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errConvert != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errConvert.Error())
	}

	price, errConvert := strconv.Atoi(req.Price)
	if errConvert != nil && errConvert.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errConvert != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errConvert.Error())
	}

	totalPrice := qty * price
	grossMargin := totalPrice - (hpp * qty)
	cleanMargin := float32(grossMargin)

	salesOrder := entity.SbsSalesOrder{
		MarketPlaceId:    req.MarketplaceId,
		OrderDate:        req.OrderDate,
		InvoiceNo:        "-",
		Sku:              req.Sku,
		ProductName:      productName,
		Qty:              qty,
		SalesPrice:       price,
		TotalPrice:       totalPrice,
		Hpp:              hpp * qty,
		GrossMargin:      grossMargin,
		PowerMerchantFee: 0,
		OngkirFee:        0,
		CleanMargin:      cleanMargin,
		Flag:             true,
		IsPayment:        true,
	}
	errResp := usecase.SbsRepository.PostSo(c, salesOrder)
	if errResp != nil && errResp.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errResp != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errResp.Error())
	}

	if req.Sku != "" {
		errUpdate := usecase.SbsRepository.DeductSbsProduct(c, req.Sku, qty)
		if errUpdate != nil && errUpdate.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errUpdate != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errUpdate.Error())
		}
	}

	return response.BuildSuccessResponse(nil)

}
