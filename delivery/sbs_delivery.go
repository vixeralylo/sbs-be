package delivery

import (
	"net/http"

	"sbs-be/model/response"
	"sbs-be/usecase"

	"github.com/gin-gonic/gin"
)

type SbsDelivery interface {
	GetSbsProduct(c *gin.Context)
	PostSo(c *gin.Context)
	NoRoute(c *gin.Context)
}

type sbsDelivery struct {
	SbsUsecase usecase.SbsUsecase
}

func GetSbsDelivery(sbsUsecase usecase.SbsUsecase) SbsDelivery {
	return &sbsDelivery{
		SbsUsecase: sbsUsecase,
	}
}

func (delivery *sbsDelivery) NoRoute(c *gin.Context) {
	resp := response.BuildRouteNotFoundResponse()
	c.JSON(http.StatusNotFound, resp)
}
