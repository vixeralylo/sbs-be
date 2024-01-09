package usecase

import (
	"context"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) UpdateSo(c context.Context, invoice_no string, status string) *response.ResponseContainer {

	if status == "Pay" {
		err := usecase.SbsRepository.UpdateSo(c, invoice_no)

		if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if err != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
		}
	} else {
		result, err := usecase.SbsRepository.GetSoById(c, invoice_no)
		if err != nil && err.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if err != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
		}

		if result[0].IsPayment {
			return response.BuildDataNotFoundResponse()
		}

		for _, salesOrder := range result {
			sku := salesOrder.Sku
			qty := salesOrder.Qty
			invoice_no := salesOrder.InvoiceNo

			errAddStock := usecase.SbsRepository.AddSbsProduct(c, sku, qty)
			if errAddStock != nil && errAddStock.Error() == config.ErrRecordNotFound.Error() {
				return response.BuildDataNotFoundResponse()
			} else if err != nil {
				return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errAddStock.Error())
			}

			errUpdateSo := usecase.SbsRepository.UpdateSo(c, invoice_no)
			if errUpdateSo != nil && errUpdateSo.Error() == config.ErrRecordNotFound.Error() {
				return response.BuildDataNotFoundResponse()
			} else if errUpdateSo != nil {
				return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errUpdateSo.Error())
			}
		}

	}

	return response.BuildSuccessResponse(nil)

}
