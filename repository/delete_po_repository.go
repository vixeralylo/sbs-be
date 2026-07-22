package repository

import (
	"context"
	"sbs-be/model/entity"
)

func (repository *sbsRepository) DeletePo(c context.Context, poNo string, sku string) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var purchaseOrder entity.SbsPurchaseOrder

	err := repository.mysqlConn.Where("po_number = ?", poNo).Where("sku = ?", sku).Delete(&purchaseOrder).Error
	if err != nil {
		return err
	}

	return nil
}
