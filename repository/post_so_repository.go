package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) PostSo(c context.Context, salesOrder entity.SbsSalesOrder) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	err := repository.mysqlConn.Create(&salesOrder).Error
	if err != nil {
		return err
	}

	return nil
}
