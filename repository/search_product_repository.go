package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) SearchProduct(c context.Context, text string) ([]entity.SbsProduct, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsProduct

	err := repository.mysqlConn.
		Where("(product_name LIKE ? OR sku LIKE ?)", "%"+text+"%", "%"+text+"%").
		Where("is_deleted = 0 OR is_deleted IS NULL").
		Order("seq asc").
		Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
