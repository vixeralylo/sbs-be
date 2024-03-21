package usecase

import (
	"context"
	"fmt"
	"time"

	"sbs-be/config"
	"sbs-be/model/constant"
	"sbs-be/model/entity"
	"sbs-be/model/response"
)

func (usecase *sbsUsecase) GetSummary(c context.Context) *response.ResponseContainer {

	// Get the current year
	currentYear := time.Now().Year()
	var sbsSummary []entity.SbsSummary

	for month := 1; month <= 12; month++ {
		monthYear := fmt.Sprintf("%d%02d", currentYear, month)

		//GET SUMMMARY SO
		resultSummarySo, errSummarySo := usecase.SbsRepository.GetSummarySo(c, monthYear)
		if errSummarySo != nil && errSummarySo.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummarySo != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummarySo.Error())
		}

		//GET SUMMMARY GAJI
		resultSummaryGaji, errSummaryGaji := usecase.SbsRepository.GetSummaryCost(c, monthYear, "Gaji")
		if errSummaryGaji != nil && errSummaryGaji.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummaryGaji != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummaryGaji.Error())
		}

		//GET SUMMMARY PLN
		resultSummaryPln, errSummaryPln := usecase.SbsRepository.GetSummaryCost(c, monthYear, "Pln")
		if errSummaryPln != nil && errSummaryPln.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummaryPln != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummaryPln.Error())
		}

		//GET SUMMMARY ADS
		resultSummaryAds, errSummaryAds := usecase.SbsRepository.GetSummaryCost(c, monthYear, "Ads")
		if errSummaryAds != nil && errSummaryAds.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummaryAds != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummaryAds.Error())
		}

		//GET SUMMMARY MATERIAL
		resultSummaryMaterial, errSummaryMaterial := usecase.SbsRepository.GetSummaryCost(c, monthYear, "Material")
		if errSummaryMaterial != nil && errSummaryMaterial.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummaryMaterial != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummaryMaterial.Error())
		}

		//GET SUMMMARY TAKE PROFIT
		resultSummaryProfit, errSummaryProfit := usecase.SbsRepository.GetSummaryCost(c, monthYear, "Profit")
		if errSummaryProfit != nil && errSummaryProfit.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummaryProfit != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummaryProfit.Error())
		}

		//GET SUMMMARY LOSS
		resultSummaryLoss, errSummaryLoss := usecase.SbsRepository.GetSummaryCost(c, monthYear, "Loss")
		if errSummaryLoss != nil && errSummaryLoss.Error() == config.ErrRecordNotFound.Error() {
			return response.BuildDataNotFoundResponse()
		} else if errSummaryLoss != nil {
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, errSummaryLoss.Error())
		}

		// Create a new SbsSummary instance for each month
		t, err := time.Parse("20060102", monthYear+"01")
		if err != nil {
			fmt.Println("Error parsing input:", err)
			return response.BuildInternalErrorResponse(constant.ERROR_CODE_DATABASE_ERROR, constant.RESPONSE_CODE_INTERNAL_ERROR, constant.RESPONSE_MESSAGE_DATABASE_ERROR, err.Error())
		}

		sbsSummaries := entity.SbsSummary{
			MonthYear: t.Format("January 2006"),
			SummaryDetail: entity.SbsSummaryDetail{
				TotalMarginSo:   resultSummarySo,
				GajiKaryawan:    resultSummaryGaji,
				Pln:             resultSummaryPln,
				TotalAds:        resultSummaryAds,
				TotalCost:       resultSummaryMaterial,
				TotalTakeProfit: resultSummaryProfit,
				TotalLoss:       resultSummaryLoss,
				SumTotal:        resultSummarySo - float32(resultSummaryGaji) - float32(resultSummaryPln) - float32(resultSummaryAds) - float32(resultSummaryMaterial) - float32(resultSummaryProfit) - float32(resultSummaryLoss),
			},
		}
		// Append the SbsSummary to the slice
		sbsSummary = append(sbsSummary, sbsSummaries)

	}

	return response.BuildSuccessResponse(sbsSummary)

}
