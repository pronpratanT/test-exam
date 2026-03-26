package models

type Test5 struct {
	ID        int    `gorm:"column:id" json:"id"`
	QueNumber string `gorm:"column:que_number;unique" json:"que_number"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
}

func (Test5) TableName() string {
	return "test5"
}
