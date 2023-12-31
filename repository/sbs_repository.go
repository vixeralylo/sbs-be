package repository

import (
	"context"
	"sbs-be/model/dto"
	"sbs-be/model/entity"

	"gorm.io/gorm"
)

type SbsRepository interface {
	Ping() string
	GetSbsProduct(c context.Context) ([]entity.SbsProduct, error)
	GetSbsProductById(c context.Context, sku string) ([]entity.SbsProduct, error)
	DeductSbsProduct(c context.Context, sku string, qty int) error
	AddSbsProduct(c context.Context, sku string, qty int) error
	GetSo(c context.Context, filter dto.RequestSo) ([]entity.SbsSalesOrder, error)
	GetSoById(c context.Context, orderId string) ([]entity.SbsSalesOrder, error)
	UpdateSo(c context.Context, invoiceNo string) error
	UpdateSoFlag(c context.Context, invoiceNo []string) error
	PostSo(c context.Context, salesOrder entity.SbsSalesOrder) error
	PostPo(c context.Context, purchasesOrder entity.SbsPurchaseOrder) error
	GetSearchProduct(c context.Context, str string) ([]entity.SbsProduct, error)
}

type sbsRepository struct {
	mysqlConn *gorm.DB
}

func GetSbsRepository(mysqlConn *gorm.DB) SbsRepository {
	return &sbsRepository{
		mysqlConn: mysqlConn,
	}
}

func (repository *sbsRepository) Ping() string {
	return "pong"
}
