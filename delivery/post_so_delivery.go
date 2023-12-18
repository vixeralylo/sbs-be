package delivery

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sbs-be/model/constant"
	"sbs-be/model/dto"
	"sbs-be/model/response"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostSo(c *gin.Context) {
	marketplace := c.GetHeader("Marketplace")
	if marketplace == "" {
		errResp := response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_HEADER_REQUEST, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_HEADER_REQ, "Tidak ada sumber marketplace")
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	if marketplace != "Tokopedia" && marketplace != "Shopee" {
		errResp := response.BuildBadRequestResponse(constant.ERROR_CODE_INVALID_HEADER_REQUEST, constant.RESPONSE_CODE_BAD_REQUEST, constant.RESPONSE_MESSAGE_INVALID_HEADER_REQ, "Marketplace harus Tokopedia atau Shopee")
		c.JSON(http.StatusBadRequest, errResp)
		return
	}

	var saleOrders []dto.RequestContainer
	var qty int
	var price int

	//TOKOPEDIA
	if marketplace == "Tokopedia" {

		xlsx, err := excelize.OpenFile("./excel/tokopedia.xlsx")
		if err != nil {
			log.Fatal("ERROR", err.Error())
		}

		sheet1Name := "Laporan Penjualan"
		rows := xlsx.GetRows(sheet1Name)

		for i := range rows {
			if i < 6 || xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i)) != "" {
				continue
			}

			qty, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", i)))
			if err != nil {
				// ... handle error
				panic(err)
			}

			price, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Q%d", i)))
			if err != nil {
				// ... handle error
				panic(err)
			}

			// Assuming columns "A" and "B" for this example
			so := dto.RequestContainer{
				InvoiceNo: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
				Sku:       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", i)),
				Qty:       qty,
				Price:     price,
			}
			saleOrders = append(saleOrders, so)
		}
	}

	//SHOPEE
	if marketplace == "Shopee" {

		xlsx, err := excelize.OpenFile("./excel/shopee.xlsx")
		if err != nil {
			log.Fatal("ERROR", err.Error())
		}

		sheet1Name := "orders"
		rows := xlsx.GetRows(sheet1Name)

		for i := range rows {
			if i < 2 || xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)) == "Batal" {
				continue
			}

			qty, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", i)))
			if err != nil {
				// ... handle error
				panic(err)
			}

			price, err = strconv.Atoi(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", i)), ""))
			if err != nil {
				// ... handle error
				panic(err)
			}

			// Assuming columns "A" and "B" for this example
			so := dto.RequestContainer{
				InvoiceNo: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
				Sku:       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("O%d", i)),
				Qty:       qty,
				Price:     price,
			}
			saleOrders = append(saleOrders, so)
		}
	}

	data := delivery.SbsUsecase.PostSo(c, marketplace, saleOrders)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
