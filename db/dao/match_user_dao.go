package dao

import (
	"badminton-go/db"
	"badminton-go/db/model"
)

const tbMatchUser = "tb_match_user"

type MatchUserDao struct {
}

func (dao *MatchUserDao) Create(uid int64, matchId int64) error {
	cli := db.Get()
	dto := &model.MatchUserDTO{
		MatchId: matchId,
		Uid:     uid,
	}
	err := cli.Table(tbMatchUser).Create(dto).Error
	return err
}

var MatchUserDaoIns = MatchUserDao{}
