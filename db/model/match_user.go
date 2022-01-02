package model

type MatchUserDTO struct {
	MatchId int64 `gorm:"column:matchId" json:"matchId"`
	Uid     int64 `gorm:"column:uid" json:"uid"`
}
