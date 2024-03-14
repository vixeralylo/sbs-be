package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) UpdatePo(c context.Context, po_no string) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var results entity.SbsPurchaseOrder
	err := repository.mysqlConn.Model(&results).Where("po_number = ?", po_no).Update("is_payment", true)
	if err != nil {
		return err.Error
	}

	return nil
}
