package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) PostPo(c context.Context, purchasesOrder entity.SbsPurchaseOrder) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	err := repository.mysqlConn.Create(&purchasesOrder).Error
	if err != nil {
		return err
	}

	return nil
}
