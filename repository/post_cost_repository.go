package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) PostCost(c context.Context, filter entity.SbsCost) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	err := repository.mysqlConn.Create(&filter).Error
	if err != nil {
		return err
	}

	return nil
}
