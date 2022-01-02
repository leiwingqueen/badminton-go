package model

import "time"

type UserDTO struct {
	Uid      int64     `gorm:"column:uid" json:"uid"`
	OpenId   string    `gorm:"column:openId" json:"openId"`
	Name     string    `gorm:"column:name" json:"name"`
	Logo     string    `gorm:"column:logo" json:"logo"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
	UpdateAt time.Time `gorm:"column:updateAt" json:"updateAt"`
}
