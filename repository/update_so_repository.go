package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) UpdateSoFlag(c context.Context, InvoiceNo []string) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var results entity.SbsSalesOrder
	err := repository.mysqlConn.Model(&results).Where("invoice_no IN ?", InvoiceNo).Update("flag", true)
	if err != nil {
		return err.Error
	}

	return nil
}
