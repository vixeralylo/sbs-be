package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostCost(c *gin.Context) {

	req := dto.RequestCost{
		Date:          c.GetHeader("costDate"),
		CostName:      c.GetHeader("costName"),
		Qty:           c.GetHeader("qty"),
		Price:         c.GetHeader("price"),
		TotalPrice:    c.GetHeader("totalPrice"),
		MarketplaceId: c.GetHeader("marketplace_id"),
		InvoiceNo:     c.GetHeader("invoice_no"),
	}

	data := delivery.SbsUsecase.PostCost(c, req)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
