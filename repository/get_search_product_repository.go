package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) GetSearchProduct(c context.Context, str string) ([]entity.SbsProduct, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsProduct

	err := repository.mysqlConn.Where("product_name like ?", "%"+str+"%").Where("is_deleted = 0 OR is_deleted IS NULL").Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
