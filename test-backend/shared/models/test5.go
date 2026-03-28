package models

import "time"

type Test5 struct {
	ID        int       `gorm:"column:id" json:"id"`
	QueNumber string    `gorm:"column:que_numbere" json:"que_number"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Iscleared bool      `gorm:"column:iscleared" json:"iscleared"`
}

func (Test5) TableName() string {
	return "test5"
}
