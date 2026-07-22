package delivery

import (
	"fmt"
	"net/http"
	"regexp"
	"sbs-be/model/dto"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostSo(c *gin.Context) {

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
	if strings.Contains(file.Filename, "Semua pesanan") {
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

	var sheetName string
	if marketplace == "Tokopedia" {
		sheetName = "OrderSKUList"
	} else {
		sheetName = "orders"
	}

	// Baca sel per-koordinat, tahan terhadap struktur XML tidak standar (mis. export TikTok)
	cells, maxRow, err := readSheetCells(uploadPath, sheetName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read excel: " + err.Error()})
		return
	}

	get := func(col string, row int) string {
		return strings.TrimSpace(cells[fmt.Sprintf("%s%d", col, row)])
	}

	var saleOrders []dto.RequestContainer

	//TOKOPEDIA
	if marketplace == "Tokopedia" {
		// baris 1 = header kolom, baris 2 = keterangan kolom, data mulai baris 3
		for r := 3; r <= maxRow; r++ {
			invoice := get("A", r)
			status := get("B", r)
			if invoice == "" || status == "Dibatalkan" {
				continue
			}

			qty, err1 := strconv.Atoi(get("J", r))
			price, err2 := strconv.Atoi(get("L", r))
			orderDate := get("AD", r)
			if err1 != nil || err2 != nil || len(orderDate) < 10 {
				continue
			}

			isPayment := true
			if status == "Belum dibayar" {
				isPayment = false
			}

			so := dto.RequestContainer{
				OrderDate: orderDate[6:10] + "-" + orderDate[3:5] + "-" + orderDate[0:2],
				InvoiceNo: invoice,
				Sku:       get("G", r),
				Qty:       qty,
				Price:     price,
				IsPayment: isPayment,
			}
			saleOrders = append(saleOrders, so)
		}
	}

	//SHOPEE
	if marketplace == "Shopee" {
		// baris 1 = header, data mulai baris 2
		for r := 2; r <= maxRow; r++ {
			invoice := get("A", r)
			status := get("B", r)
			if invoice == "" || status == "Batal" {
				continue
			}

			qty, err1 := strconv.Atoi(get("S", r))
			priceClean := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(get("R", r), "")
			price, err2 := strconv.Atoi(priceClean)
			orderDate := get("J", r)
			if err1 != nil || err2 != nil || len(orderDate) < 10 {
				continue
			}

			isPayment := true
			if status == "Belum Bayar" {
				isPayment = false
			}

			so := dto.RequestContainer{
				OrderDate: orderDate[0:10],
				InvoiceNo: invoice,
				Sku:       get("O", r),
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
