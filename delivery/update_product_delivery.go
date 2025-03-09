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

	data := delivery.SbsUsecase.UpdateSbsProduct(c, sku, qty, hppConvert, priceConvert)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
