package model

type RoundDTO struct {
	MatchId int64  `gorm:"column:matchId" json:"id"`
	RoundId int64  `gorm:"column:roundId" json:"roundId"`
	Detail  string `gorm:"column:detail" json:"detail"`
	Result  string `gorm:"column:detail" json:"detail"`
}
