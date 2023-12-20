package repository

import (
	"context"

	"sbs-be/model/entity"

	"gorm.io/gorm"
)

func (repository *sbsRepository) DeductSbsProduct(c context.Context, sku string, qty int) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var results entity.SbsProduct

	err := repository.mysqlConn.Model(&results).Where("sku = ?", sku).UpdateColumn("stock", gorm.Expr("stock - ?", qty)).Error
	if err != nil {
		return err
	}

	return nil
}

func (repository *sbsRepository) AddSbsProduct(c context.Context, sku string, qty int) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var results entity.SbsProduct

	err := repository.mysqlConn.Model(&results).Where("sku = ?", sku).UpdateColumn("stock", gorm.Expr("stock + ?", qty)).Error
	if err != nil {
		return err
	}

	return nil
}
