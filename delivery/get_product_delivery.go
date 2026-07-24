package delivery

import (
	"net/http"

	"sbs-be/model/dto"

	"github.com/gin-gonic/gin"
)

func (delivery *sbsDelivery) GetSbsProduct(c *gin.Context) {
	// request body opsional; jika ada is_deleted=true tampilkan produk nonaktif
	var reqBody dto.RequestBody
	_ = c.ShouldBindJSON(&reqBody)

	data := delivery.SbsUsecase.GetSbsProduct(c, reqBody.IsDeleted)

	if data.StatusCode >= 400 && data.StatusCode != http.StatusNotFound {
		c.JSON(data.StatusCode, data)
		return
	}

	c.JSON(http.StatusOK, data)
}
