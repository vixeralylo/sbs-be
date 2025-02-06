package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) DeletePo(c *gin.Context) {

	poNo := c.GetHeader("po_no")
	sku := c.GetHeader("sku")

	data := delivery.SbsUsecase.DeletePo(c, poNo, sku)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
