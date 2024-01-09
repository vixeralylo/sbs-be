package app

import (
	"sbs-be/delivery"
	"sbs-be/middleware"
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

	router.Use(middleware.CORSMiddleware())

	svc := router.Group("/")
	svc.GET("/product", sbsDelivery.GetSbsProduct)
	svc.GET("/so", sbsDelivery.GetSo)
	svc.POST("/so", sbsDelivery.PostSo)
	svc.POST("/po", sbsDelivery.PostPo)
	svc.PUT("/so", sbsDelivery.UpdateSo)
	svc.GET("/search", sbsDelivery.GetSearchProduct)

	router.NoRoute(sbsDelivery.NoRoute)

	return router
}
