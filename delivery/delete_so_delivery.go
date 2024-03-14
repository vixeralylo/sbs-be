package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) DeleteSo(c *gin.Context) {

	invoiceNo := c.GetHeader("invoice_no")

	data := delivery.SbsUsecase.DeleteSo(c, invoiceNo)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
