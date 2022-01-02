package model

import "time"

type MatchDTO struct {
	MatchId  int64     `gorm:"column:matchId" json:"id"`
	Name     string    `gorm:"column:name" json:"name"`
	StartAt  time.Time `gorm:"column:startAt" json:"startAt"`
	CreateAt time.Time `gorm:"column:createAt" json:"createAt"`
	UpdateAt time.Time `gorm:"column:updateAt" json:"updateAt"`
	Status   int32     `gorm:"column:status" json:"status"`
}
