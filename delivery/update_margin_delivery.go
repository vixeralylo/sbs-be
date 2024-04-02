package delivery

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sbs-be/model/dto"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) UpdateMargin(c *gin.Context) {

	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to parse form data"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
		return
	}

	var marketplace string
	if strings.Contains(file.Filename, "Tokopedia") {
		marketplace = "Tokopedia"
	} else {
		marketplace = "Shopee"
	}

	// Save the uploaded file to a temporary location
	uploadPath := "excel/" + file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	var saleOrders []dto.RequestContainer
	var qty int
	var price int

	xlsx, err := excelize.OpenFile(uploadPath)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	//TOKOPEDIA
	if marketplace == "Tokopedia" {
		sheet1Name := "Laporan Penjualan"
		rows := xlsx.GetRows(sheet1Name)

		for i := range rows {
			if i+1 < 6 || (xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i+1)) != "" && xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i+1)) != "Nama Toko: ") {
				continue
			}

			qty, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("N%d", i+1)))
			if err != nil {
				// ... handle error
				panic(err)
			}

			price, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("Q%d", i+1)))
			if err != nil {
				// ... handle error
				panic(err)
			}

			isPayment := true
			if xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1)) == "Belum Bayar" {
				isPayment = false
			}

			// Assuming columns "A" and "B" for this example
			so := dto.RequestContainer{
				OrderDate: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i+1))[6:10] + "-" + xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i+1))[3:5] + "-" + xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i+1))[0:2],
				InvoiceNo: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1)),
				Sku:       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("K%d", i+1)),
				Qty:       qty,
				Price:     price,
				IsPayment: isPayment,
			}
			saleOrders = append(saleOrders, so)
		}
	}

	//SHOPEE
	if marketplace == "Shopee" {
		sheet1Name := "orders"
		rows := xlsx.GetRows(sheet1Name)

		for i := range rows {
			if i+1 < 2 || xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1)) == "Batal" {
				continue
			}

			qty, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("S%d", i+1)))
			if err != nil {
				// ... handle error
				panic(err)
			}

			price, err = strconv.Atoi(regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("R%d", i+1)), ""))
			if err != nil {
				// ... handle error
				panic(err)
			}

			isPayment := true
			if xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1)) == "Belum Bayar" {
				isPayment = false
			}

			// Assuming columns "A" and "B" for this example
			so := dto.RequestContainer{
				OrderDate: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("J%d", i+1))[0:10],
				InvoiceNo: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i+1)),
				Sku:       xlsx.GetCellValue(sheet1Name, fmt.Sprintf("O%d", i+1)),
				Qty:       qty,
				Price:     price,
				IsPayment: isPayment,
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
