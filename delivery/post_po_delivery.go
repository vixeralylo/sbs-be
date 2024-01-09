package delivery

import (
	"fmt"
	"log"
	"net/http"
	"sbs-be/model/dto"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostPo(c *gin.Context) {

	var purchaseOrder []dto.RequestPo
	var qty int
	var price float64
	var discount float64
	var ppn float64
	var total_price float64

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

	// Save the uploaded file to a temporary location
	uploadPath := "excel/" + file.Filename
	if err := c.SaveUploadedFile(file, uploadPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	xlsx, err := excelize.OpenFile(uploadPath)
	if err != nil {
		log.Fatal("ERROR", err.Error())
	}

	sheet1Name := "Sheet1"
	rows := xlsx.GetRows(sheet1Name)

	for i := range rows {
		if i+1 < 2 {
			continue
		}

		qty, err = strconv.Atoi(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("E%d", i+1)))
		if err != nil {
			// ... handle error
			panic(err)
		}

		price, err = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("F%d", i+1)), 32)
		if err != nil {
			// ... handle error
			panic(err)
		}

		discount, err = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("G%d", i+1)), 32)
		if err != nil {
			// ... handle error
			panic(err)
		}

		ppn, err = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("H%d", i+1)), 32)
		if err != nil {
			// ... handle error
			panic(err)
		}

		total_price, err = strconv.ParseFloat(xlsx.GetCellValue(sheet1Name, fmt.Sprintf("I%d", i+1)), 32)
		if err != nil {
			// ... handle error
			panic(err)
		}

		// Assuming columns "A" and "B" for this example
		po := dto.RequestPo{
			PoNumber:    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i+1)),
			PoDate:      xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i+1)),
			Sku:         xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i+1)),
			ProductName: xlsx.GetCellValue(sheet1Name, fmt.Sprintf("D%d", i+1)),
			Qty:         qty,
			Price:       price,
			Discount:    discount,
			Ppn:         ppn,
			TotalPrice:  total_price,
		}
		purchaseOrder = append(purchaseOrder, po)
	}

	data := delivery.SbsUsecase.PostPo(c, purchaseOrder)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
