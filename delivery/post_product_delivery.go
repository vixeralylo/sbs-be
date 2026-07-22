package delivery

import (
	"net/http"

	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) PostProduct(c *gin.Context) {

	var reqBody dto.RequestAddProduct
	// Bind JSON request body to struct
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	data := delivery.SbsUsecase.PostProduct(c, reqBody)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
