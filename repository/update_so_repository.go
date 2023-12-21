package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) UpdateSo(c context.Context, invoice_no string) error {

	if c.Err() == context.DeadlineExceeded {
		return c.Err()
	}

	var results entity.SbsSalesOrder
	err := repository.mysqlConn.Model(&results).Where("invoice_no = ?", invoice_no).Update("is_payment", true)
	if err != nil {
		return err.Error
	}

	return nil
}

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
