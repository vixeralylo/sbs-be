package delivery

import (
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetSo(c *gin.Context) {

	marketplace_id := c.GetHeader("marketplace_id")
	start_date := c.GetHeader("start_date")
	end_date := c.GetHeader("end_date")
	invoice_no := c.GetHeader("invoice_no")
	is_not_payment := c.GetHeader("is_not_payment")

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
