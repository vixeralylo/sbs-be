package repository

import (
	"context"
	"sbs-be/model/entity"

	"gorm.io/gorm"
)

type SbsRepository interface {
	Ping() string
	GetSbsProduct(c context.Context) ([]entity.SbsProduct, error)
	GetSbsProductById(c context.Context, sku string) ([]entity.SbsProduct, error)
	UpdateSbsProduct(c context.Context, sku string, qty int) error
	PostSo(c context.Context, salesOrder entity.SbsSalesOrder) error
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
