package models

import "time"

type Test4 struct {
	ID         int       `gorm:"column:id" json:"id"`
	FName      string    `gorm:"column:fname" json:"fname"`
	LName      string    `gorm:"column:lname" json:"lname"`
	Email      string    `gorm:"column:email" json:"email"`
	Phone      string    `gorm:"column:phone" json:"phone"`
	Profile    string    `gorm:"column:profile" json:"profile"`
	Birthdate  time.Time `gorm:"column:birthdate" json:"birthdate"`
	Occupation string    `gorm:"column:occupation" json:"occupation"`
	Sex        string    `gorm:"column:sex" json:"sex"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Test4) TableName() string {
	return "test4"
}
