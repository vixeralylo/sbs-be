package repository

import (
	"context"

	"sbs-be/model/dto"
	"sbs-be/model/entity"
)

func (repository *sbsRepository) GetCost(c context.Context, filter dto.RequestCost) ([]entity.SbsCost, error) {

	if c.Err() == context.DeadlineExceeded {
		return nil, c.Err()
	}

	var results []entity.SbsCost
	var start_date = filter.StartDate
	var end_date = filter.EndDate

	dbTemp := repository.mysqlConn.Table(entity.TABLE_COST)

	if len(start_date) > 0 && len(end_date) > 0 {
		dbTemp = dbTemp.Where("date BETWEEN ? AND ?", start_date, end_date)
	}

	err := dbTemp.Find(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
