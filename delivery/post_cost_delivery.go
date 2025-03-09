package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostCost(c *gin.Context) {

	var reqBody dto.RequestBody
	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Access values
	cost_date := reqBody.CostDate
	cost_type := reqBody.CostType
	cost_name := reqBody.CostName
	qty := reqBody.Qty
	price := reqBody.Price
	added_price := reqBody.AddedPrice
	total_price := reqBody.TotalPrice
	marketplace_id := reqBody.MarketplaceId
	invoice_no := reqBody.InvoiceNo

	req := dto.RequestCost{
		Date:          cost_date,
		CostType:      cost_type,
		CostName:      cost_name,
		Qty:           qty,
		Price:         price,
		AddedPrice:    added_price,
		TotalPrice:    total_price,
		MarketplaceId: marketplace_id,
		InvoiceNo:     invoice_no,
	}

	data := delivery.SbsUsecase.PostCost(c, req)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
