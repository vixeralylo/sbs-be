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
	svc.POST("/product/get", sbsDelivery.GetSbsProduct)
	svc.POST("/product/put", sbsDelivery.UpdateSbsProduct)
	svc.POST("/so/get", sbsDelivery.GetSo)
	svc.POST("/so/post", sbsDelivery.PostSo)
	svc.POST("/so_manual/post", sbsDelivery.PostSoManual)
	svc.POST("/so/put", sbsDelivery.UpdateSo)
	svc.POST("/so/delete", sbsDelivery.DeleteSo)
	svc.POST("/po/get", sbsDelivery.GetPo)
	svc.POST("/po/delete", sbsDelivery.DeletePo)
	svc.POST("/po/post", sbsDelivery.PostPo)
	svc.POST("/po/put", sbsDelivery.UpdatePo)
	svc.POST("/search/get", sbsDelivery.GetSearchProduct)
	svc.POST("/cost/get", sbsDelivery.GetCost)
	svc.POST("/cost/post", sbsDelivery.PostCost)
	svc.POST("/summary/get", sbsDelivery.GetSummary)
	svc.POST("/margin/put", sbsDelivery.UpdateMargin)

	router.NoRoute(sbsDelivery.NoRoute)

	return router
}
