package common

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo
	Debug bool
	Likes []string // Fuzzy matching field
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	if option.Sort == "" {
		option.Sort = "created_at desc" // Default sort by time
	}

	DB = DB.Where(model)

	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}

		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	count = DB.Find(&list).RowsAffected

	query := DB.Where(model)

	//Default returns all values
	if option.Limit == 0 {
		return list, count, err
	}

	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
