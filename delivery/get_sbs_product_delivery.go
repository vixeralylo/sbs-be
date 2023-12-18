package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetSbsProduct(c *gin.Context) {
	data := delivery.SbsUsecase.GetSbsProduct(c)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
