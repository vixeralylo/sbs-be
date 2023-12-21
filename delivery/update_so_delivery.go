package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) UpdateSo(c *gin.Context) {

	invoiceNo := c.GetHeader("invoice_no")
	status := c.GetHeader("status")

	data := delivery.SbsUsecase.UpdateSo(c, invoiceNo, status)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
