package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) PostProduct(c context.Context, product entity.SbsProduct) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	err := repository.mysqlConn.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}
