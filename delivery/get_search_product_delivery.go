package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetSearchProduct(c *gin.Context) {
	str := c.GetHeader("string")

	data := delivery.SbsUsecase.GetSearchProduct(c, str)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
