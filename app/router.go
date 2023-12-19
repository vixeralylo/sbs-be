package app

import (
	"sbs-be/delivery"
	"sbs-be/repository"
	"sbs-be/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(
	mysqlConn *gorm.DB,
) *gin.Engine {

	sbsRepository := repository.GetSbsRepository(mysqlConn)
	sbsUsecase := usecase.GetSbsUsecase(sbsRepository)
	sbsDelivery := delivery.GetSbsDelivery(sbsUsecase)

	router := gin.Default()

	svc := router.Group("/")
	svc.GET("/product", sbsDelivery.GetSbsProduct)
	svc.GET("/so", sbsDelivery.GetSo)
	svc.POST("/so", sbsDelivery.PostSo)
	svc.POST("/po", sbsDelivery.PostPo)

	router.NoRoute(sbsDelivery.NoRoute)

	return router
}
