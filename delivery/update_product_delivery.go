package delivery

import (
	"fmt"
	"net/http"
	"sbs-be/model/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) UpdateSbsProduct(c *gin.Context) {

	var reqBody dto.RequestBody
	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Access values
	hpp := reqBody.Hpp
	price := reqBody.Price
	sku := reqBody.Sku
	qty := reqBody.Qty
	seq := reqBody.Seq
	adminFee := reqBody.AdminFee
	ongkirFee := reqBody.OngkirFee
	tax := reqBody.Tax
	isDeleted := reqBody.IsDeleted

	// Convert string to float32
	hppConvert, err := strconv.ParseFloat(hpp, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert string to float32
	priceConvert, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert string to float64 (default 0 kalau kosong)
	adminFeeConvert := 0.0
	if adminFee != "" {
		adminFeeConvert, err = strconv.ParseFloat(adminFee, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	ongkirFeeConvert := 0.0
	if ongkirFee != "" {
		ongkirFeeConvert, err = strconv.ParseFloat(ongkirFee, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	taxConvert := 0.0
	if tax != "" {
		taxConvert, err = strconv.ParseFloat(tax, 64)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	data := delivery.SbsUsecase.UpdateSbsProduct(c, sku, qty, seq, hppConvert, priceConvert, adminFeeConvert, ongkirFeeConvert, taxConvert, isDeleted)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
