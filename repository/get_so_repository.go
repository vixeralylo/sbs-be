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

	dbTemp := repository.mysqlConn.Table(entity.TABLE_SALES_ORDER)

	if len(marketplace_id) > 0 {
		dbTemp = dbTemp.Where("marketplace_id = ?", marketplace_id)
	}

	if len(start_date) > 0 && len(end_date) > 0 {
		dbTemp = dbTemp.Where("order_date BETWEEN ? AND ?", start_date, end_date)
	}

	err := dbTemp.Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
