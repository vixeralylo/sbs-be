package delivery

import (
	"log"
	"net/http"
	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetPo(c *gin.Context) {
	var reqBody dto.RequestBody

	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Access values
	start_date := reqBody.StartDate
	end_date := reqBody.EndDate
	is_not_payment := reqBody.IsNotPayment

	log.Println("start_date:", start_date)
	log.Println("end_date:", end_date)
	log.Println("is_not_payment:", is_not_payment)

	filter := dto.RequestPo{
		StartDate:    start_date,
		EndDate:      end_date,
		IsNotPayment: is_not_payment,
	}

	data := delivery.SbsUsecase.GetPo(c, filter)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
