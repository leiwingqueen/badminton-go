package dao

import (
	"badminton-go/db"
	"badminton-go/db/model"
)

const tbMatch = "tb_match"

type MatchDao struct {
}

func (dao *MatchDao) List(page int, size int) ([]model.CounterModel, error) {
	var err error
	cli := db.Get()
	offset := (page - 1) * size
	var res []model.CounterModel
	err = cli.Table(tbMatch).Order("createAt desc").Offset(offset).Limit(size).Find(&res).Error
	return res, err
}

var MatchDaoIns = MatchDao{}
