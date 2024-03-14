package repository

import (
	"context"

	"sbs-be/model/dto"
	"sbs-be/model/entity"
)

func (repository *sbsRepository) GetSo(c context.Context, filter dto.RequestSo) ([]entity.SbsSalesOrder, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsSalesOrder
	var marketplace_id = filter.MarketplaceId
	var start_date = filter.StartDate
	var end_date = filter.EndDate
	var so_number = filter.SoNumber

	dbTemp := repository.mysqlConn.Table(entity.TABLE_SALES_ORDER)

	if len(marketplace_id) > 0 {
		dbTemp = dbTemp.Where("marketplace_id = ?", marketplace_id)
	}

	if len(start_date) > 0 && len(end_date) > 0 {
		dbTemp = dbTemp.Where("order_date BETWEEN ? AND ?", start_date, end_date)
	}

	if len(so_number) > 0 {
		dbTemp = dbTemp.Where("invoice_no = ?", so_number)
	}

	err := dbTemp.Where("is_cancel = ?", false).Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func (repository *sbsRepository) GetSoById(c context.Context, orderId string) ([]entity.SbsSalesOrder, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsSalesOrder

	err := repository.mysqlConn.Where("invoice_no = ?", orderId).Where("flag is true").Where("is_cancel = ?", false).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
