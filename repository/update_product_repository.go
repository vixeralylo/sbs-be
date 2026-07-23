package repository

import (
	"context"

	"sbs-be/model/entity"

	"gorm.io/gorm"
)

func (repository *sbsRepository) UpdateSbsProduct(c context.Context, sku string, qty int, seq int, hpp float64, price float64, adminFee float64, ongkirFee float64, tax float64) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var results entity.SbsProduct

	err := repository.mysqlConn.Model(&results).Where("sku = ?", sku).
		Updates(map[string]interface{}{
			"stock":      qty,
			"hpp":        hpp,
			"price":      price,
			"seq":        seq,
			"admin_fee":  adminFee,
			"ongkir_fee": ongkirFee,
			"tax":        tax,
		}).Error
	if err != nil {
		return err
	}

	return nil
}

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
