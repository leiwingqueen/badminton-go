package model

import "time"

type MatchDTO struct {
	MatchId   int64     `gorm:"column:matchId" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	StartAt   time.Time `gorm:"column:startAt" json:"startAt"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	Status    int32     `gorm:"column:status" json:"status"`
}
