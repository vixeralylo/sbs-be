package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) GetSbsProduct(c context.Context) ([]entity.SbsProduct, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsProduct

	err := repository.mysqlConn.Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repository *sbsRepository) GetSbsProductById(c context.Context, sku string) ([]entity.SbsProduct, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsProduct

	err := repository.mysqlConn.Where("sku = ?", sku).Find(&results).Limit(1).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
