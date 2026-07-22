package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetSo(c *gin.Context) {

	var reqBody dto.RequestBody
	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	marketplace_id := reqBody.MarketplaceId
	start_date := reqBody.StartDate
	end_date := reqBody.EndDate
	is_not_payment := reqBody.IsNotPayment
	invoice_no := reqBody.InvoiceNo

	filter := dto.RequestSo{
		MarketplaceId: marketplace_id,
		StartDate:     start_date,
		EndDate:       end_date,
		SoNumber:      invoice_no,
		IsNotPayment:  is_not_payment,
	}

	data := delivery.SbsUsecase.GetSo(c, filter)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
