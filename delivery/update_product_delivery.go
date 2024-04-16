package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) UpdateSbsProduct(c *gin.Context) {

	sku := c.GetHeader("sku")
	qty := c.GetHeader("qty")

	data := delivery.SbsUsecase.UpdateSbsProduct(c, sku, qty)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
