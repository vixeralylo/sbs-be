package repository

import (
	"context"

	"sbs-be/model/dto"
	"sbs-be/model/entity"
)

func (repository *sbsRepository) GetPo(c context.Context, filter dto.RequestPo) ([]entity.SbsPurchaseOrder, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsPurchaseOrder
	var start_date = filter.StartDate
	var end_date = filter.EndDate
	var is_not_payment = filter.IsNotPayment

	dbTemp := repository.mysqlConn.Table(entity.TABLE_PURCHASE_ORDER)

	if len(start_date) > 0 && len(end_date) > 0 {
		dbTemp = dbTemp.Where("po_date BETWEEN ? AND ?", start_date, end_date)
	}

	if len(is_not_payment) > 0 {

		var isNotPayment bool
		if is_not_payment == "true" {
			isNotPayment = true
		} else {
			isNotPayment = false
		}

		dbTemp = dbTemp.Where("is_payment = ?", isNotPayment)
	}

	err := dbTemp.Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
