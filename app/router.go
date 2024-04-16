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

	svc := router.Group("/api")
	svc.GET("/product", sbsDelivery.GetSbsProduct)
	svc.PUT("/product", sbsDelivery.UpdateSbsProduct)
	svc.GET("/so", sbsDelivery.GetSo)
	svc.POST("/so", sbsDelivery.PostSo)
	svc.POST("/so_manual", sbsDelivery.PostSoManual)
	svc.PUT("/so", sbsDelivery.UpdateSo)
	svc.DELETE("/so", sbsDelivery.DeleteSo)
	svc.GET("/po", sbsDelivery.GetPo)
	svc.POST("/po", sbsDelivery.PostPo)
	svc.PUT("/po", sbsDelivery.UpdatePo)
	svc.GET("/search", sbsDelivery.GetSearchProduct)
	svc.GET("/cost", sbsDelivery.GetCost)
	svc.POST("/cost", sbsDelivery.PostCost)
	svc.GET("/summary", sbsDelivery.GetSummary)
	svc.PUT("/margin", sbsDelivery.UpdateMargin)

	router.NoRoute(sbsDelivery.NoRoute)

	return router
}
