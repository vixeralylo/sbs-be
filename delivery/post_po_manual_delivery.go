package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostPoManual(c *gin.Context) {

	var purchaseOrder []dto.RequestPo
	// Bind JSON request body to struct (array of PO line items)
	if err := c.ShouldBindJSON(&purchaseOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if len(purchaseOrder) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty purchase order"})
		return
	}

	data := delivery.SbsUsecase.PostPo(c, purchaseOrder)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
