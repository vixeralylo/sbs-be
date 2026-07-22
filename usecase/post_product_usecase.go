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

func (usecase *sbsUsecase) PostProduct(c context.Context, req dto.RequestAddProduct) *response.ResponseContainer {

	// SKU wajib diisi
	if req.Sku == "" {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_BODY_REQUEST, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_BODY_REQ, "sku is required")
	}

	// Cek apakah SKU sudah ada
	existing, _ := usecase.SbsRepository.GetSbsProductById(c, req.Sku)
	if len(existing) > 0 {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_ALREADY_EXIST, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_ALREADY_EXIST, "sku already exists")
	}

	hpp, err := strconv.Atoi(req.Hpp)
	if err != nil {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid hpp")
	}

	price, err := strconv.ParseFloat(req.Price, 64)
	if err != nil {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid price")
	}

	adminFee, err := strconv.ParseFloat(req.AdminFee, 64)
	if err != nil {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid admin_fee")
	}

	ongkirFee, err := strconv.ParseFloat(req.OngkirFee, 64)
	if err != nil {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid ongkir_fee")
	}

	tax, err := strconv.ParseFloat(req.Tax, 64)
	if err != nil {
		return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid tax")
	}

	// seq opsional, default 0 kalau kosong
	seq := 0
	if req.Seq != "" {
		seq, err = strconv.Atoi(req.Seq)
		if err != nil {
			return response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_DATA_TYPE, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_DATA_TYPE, "invalid seq")
		}
	}

	product := entity.SbsProduct{
		Sku:         req.Sku,
		ProductName: req.ProductName,
		Stock:       req.Qty,
		Hpp:         hpp,
		Price:       price,
		AdminFee:    adminFee,
		OngkirFee:   ongkirFee,
		Tax:         tax,
		Seq:         seq,
	}

	errInsert := usecase.SbsRepository.PostProduct(c, product)
	if errInsert != nil && errInsert.Error() == config.ErrRecordNotFound.Error() {
		return response.BuildDataNotFoundResponse()
	} else if errInsert != nil {
		return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errInsert.Error())
	}

	return response.BuildSuccessResponse(nil)
}
