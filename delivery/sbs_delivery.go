package delivery

import (
	"net/http"

	"sbs-be/model/response"
	"sbs-be/usecase"

	"github.com/gin-gonic/gin"
)

type SbsDelivery interface {
	GetSbsProduct(c *gin.Context)
	GetSo(c *gin.Context)
	PostSo(c *gin.Context)
	PostSoManual(c *gin.Context)
	GetPo(c *gin.Context)
	PostPo(c *gin.Context)
	UpdateSo(c *gin.Context)
	GetSearchProduct(c *gin.Context)
	DeleteSo(c *gin.Context)
	UpdatePo(c *gin.Context)
	GetCost(c *gin.Context)
	PostCost(c *gin.Context)
	GetSummary(c *gin.Context)
	UpdateMargin(c *gin.Context)
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
