package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostSoManual(c *gin.Context) {

	req := dto.RequestSoManual{
		MarketplaceId: c.GetHeader("marketplace_id"),
		OrderDate:     c.GetHeader("soDate"),
		Sku:           c.GetHeader("sku"),
		Qty:           c.GetHeader("qty"),
		Price:         c.GetHeader("price"),
		TotalPrice:    c.GetHeader("total_price"),
		SoNumber:      c.GetHeader("soNumber"),
	}

	data := delivery.SbsUsecase.PostSoManual(c, req)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
