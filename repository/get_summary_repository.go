package repository

import (
	"context"

	"sbs-be/model/entity"
)

func (repository *sbsRepository) GetSummarySo(c context.Context, monthYear string) (float32, error) {

	if c.Err() == context.DeadlineExceeded {
		return 0, c.Err()
	}

	var results entity.SbsSalesOrder
	var month = monthYear[4:]
	var year = monthYear[0:4]

	dbTemp := repository.mysqlConn.
		Table(entity.TABLE_SALES_ORDER).
		Select("SUM(clean_margin) as clean_margin").
		Where("MONTH(order_date) = ?", month).
		Where("YEAR(order_date) = ?", year)

	err := dbTemp.Find(&results).Error

	if err != nil {
		return 0, err
	}

	return results.CleanMargin, nil
}

func (repository *sbsRepository) GetSummaryCost(c context.Context, monthYear string, costType string) (int, error) {

	if c.Err() == context.DeadlineExceeded {
		return 0, c.Err()
	}

	var results entity.SbsSalesOrder
	var month = monthYear[4:]
	var year = monthYear[0:4]

	dbTemp := repository.mysqlConn.
		Table(entity.TABLE_COST).
		Select("SUM(total_price) as total_price").
		Where("MONTH(date) = ?", month).
		Where("YEAR(date) = ?", year).
		Where("cost_type = ?", costType)

	err := dbTemp.Find(&results).Error

	if err != nil {
		return 0, err
	}

	return results.TotalPrice, nil
}
