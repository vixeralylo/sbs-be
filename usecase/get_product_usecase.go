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
		admin := 0.045 * price
		ongkir := 0.04 * price
		cleanMargin := gross - admin - ongkir
		stock, _ := strconv.ParseFloat(product.Stock, 64)

		var pct float64
		if cleanMargin != 0 {
			pct = toFixed((cleanMargin/hpp)*100, 2)
		} else {
			pct = 0
		}

		sisaPersesiaan = sisaPersesiaan + (stock * hpp)

		products := entity.SbsProduct{
			Sku:         product.Sku,
			ProductName: product.ProductName,
			Stock:       product.Stock,
			Hpp:         product.Hpp,
			Price:       product.Price,
			Seq:         product.Seq,
			Gross:       gross,
			Admin:       admin,
			Ongkir:      ongkir,
			CleanMargin: cleanMargin,
			Pct:         pct,
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
