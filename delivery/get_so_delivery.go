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

	filter := dto.RequestSo{
		MarketplaceId: marketplace_id,
		StartDate:     start_date,
		EndDate:       end_date,
	}

	data := delivery.SbsUsecase.GetSo(c, filter)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
