package delivery

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) UpdateSbsProduct(c *gin.Context) {

	sku := c.GetHeader("sku")
	qty := c.GetHeader("qty")
	hpp := c.GetHeader("hpp")
	price := c.GetHeader("price")

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
