package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) UpdatePo(c *gin.Context) {

	po_no := c.GetHeader("po_no")
	status := c.GetHeader("status")

	data := delivery.SbsUsecase.UpdatePo(c, po_no, status)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
