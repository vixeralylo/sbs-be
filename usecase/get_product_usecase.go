package usecase

import (
	"context"
	"math"
	"strconv"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) GetSbsProduct(c context.Context) *response.ResponseContainer {

	resultProduct, err := usecase.SbsRepository.GetSbsProduct(c)
	var productList []entity.SbsProduct
	var sisaPersesiaan float64

	for _, product := range resultProduct {
		price := float64(product.Price)
		hpp := float64(product.Hpp)
		gross := price - hpp
		adminTok := (product.AdminFeeTok / 100) * price
		ongkirTok := (product.OngkirFeeTok / 100) * price
		adminSho := (product.AdminFeeSho / 100) * price
		ongkirSho := (product.OngkirFeeSho / 100) * price
		cleanMarginTok := gross - adminTok - ongkirTok
		cleanMarginSho := gross - adminSho - ongkirSho
		stock, _ := strconv.ParseFloat(product.Stock, 64)

		var pctTok, pctSho float64
		if cleanMarginTok != 0 {
			pctTok = toFixed((cleanMarginTok/hpp)*100, 2)
			pctSho = toFixed((cleanMarginSho/hpp)*100, 2)
		} else {
			pctTok = 0
			pctSho = 0
		}

		sisaPersesiaan = sisaPersesiaan + (stock * hpp)

		products := entity.SbsProduct{
			Sku:            product.Sku,
			ProductName:    product.ProductName,
			Stock:          product.Stock,
			Hpp:            product.Hpp,
			Price:          product.Price,
			Seq:            product.Seq,
			Gross:          gross,
			AdminFeeTok:    adminTok,
			OngkirFeeTok:   ongkirTok,
			AdminFeeSho:    adminSho,
			OngkirFeeSho:   ongkirSho,
			CleanMarginTok: cleanMarginTok,
			CleanMarginSho: cleanMarginSho,
			PctTok:         pctTok,
			PctSho:         pctSho,
		}
		productList = append(productList, products)
	}

	result := entity.SbsProductResponse{
		SisaPersesiaan: sisaPersesiaan,
		SbsProductList: productList,
	}

	if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if err != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
	}

	return response.BuildSuccessResponse(result)

}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
