package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetSummary(c *gin.Context) {

	data := delivery.SbsUsecase.GetSummary(c)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
