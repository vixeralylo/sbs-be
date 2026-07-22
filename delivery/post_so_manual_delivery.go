package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostSoManual(c *gin.Context) {
	var reqBody dto.RequestBody
	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Access values
	marketplace_id := reqBody.MarketplaceId
	so_date := reqBody.SoDate
	sku := reqBody.Sku
	qty := reqBody.Qty
	price := reqBody.Price
	total_price := reqBody.TotalPrice
	invoice_no := reqBody.InvoiceNo

	req := dto.RequestSoManual{
		MarketplaceId: marketplace_id,
		OrderDate:     so_date,
		Sku:           sku,
		Qty:           qty,
		Price:         price,
		TotalPrice:    total_price,
		SoNumber:      invoice_no,
	}

	data := delivery.SbsUsecase.PostSoManual(c, req)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
