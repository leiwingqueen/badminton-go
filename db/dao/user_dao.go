package dao

import (
	"badminton-go/db"
	"badminton-go/db/model"
	"time"
)

const tbUser string = "tb_user"

type UserDao struct {
}

func (dao *UserDao) Insert(openId string, logo string, name string) (int64, error) {
	var err error
	cli := db.Get()
	dto := &model.UserDTO{}
	now := time.Now()
	dto.Name = name
	dto.Logo = logo
	dto.OpenId = openId
	dto.CreateAt = now
	dto.UpdateAt = now
	err = cli.Table(tbUser).Save(dto).Error
	return dto.Uid, err
}

func (dao *UserDao) getByUid(uid int64) (*model.UserDTO, error) {
	var err error
	var user = new(model.UserDTO)
	cli := db.Get()
	err = cli.Table(tbUser).Where("uid = ?", uid).First(user).Error
	return user, err
}

func (dao *UserDao) getByOpenId(openId int64) (*model.UserDTO, error) {
	var err error
	var user = new(model.UserDTO)
	cli := db.Get()
	err = cli.Table(tbUser).Where("openId = ?", openId).First(user).Error
	return user, err
}
