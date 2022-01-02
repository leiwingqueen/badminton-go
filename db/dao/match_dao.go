package dao

import (
	"badminton-go/db"
	"badminton-go/db/model"
	"time"
)

const tbMatch = "tb_match"

type MatchDao struct {
}

func (dao *MatchDao) List(page int, size int) ([]model.MatchDTO, error) {
	var err error
	cli := db.Get()
	offset := (page - 1) * size
	var res []model.MatchDTO
	err = cli.Table(tbMatch).Order("createAt desc").Offset(offset).Limit(size).Find(&res).Error
	return res, err
}

func (dao *MatchDao) Create(uid int64, name string, startTime time.Time) (int64, error) {
	var err error
	cli := db.Get()
	dto := &model.MatchDTO{}
	now := time.Now()
	dto.Name = name
	dto.StartAt = startTime
	dto.CreateAt = now
	dto.UpdateAt = now
	err = cli.Table(tbMatch).Save(dto).Error
	return dto.MatchId, err
}

var MatchDaoIns = MatchDao{}
