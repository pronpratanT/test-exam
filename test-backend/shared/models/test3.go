package models

type Test3 struct {
	ID     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Reason string `gorm:"column:reason" json:"reason"`
	Status string `gorm:"column:status" json:"status"`
}

func (Test3) TableName() string {
	return "test3"
}
