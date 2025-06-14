package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
)

func (s *sqlStore) ListWithCondition(ctx context.Context, paging *common.Paging, filter *restaurantmodel.Filter) ([]restaurantmodel.Restaurant, error) {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}

		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	db = db.Order("id desc")

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrorDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	var result []restaurantmodel.Restaurant

	if err := db.
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
