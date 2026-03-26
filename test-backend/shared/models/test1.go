package models

type Test1 struct {
	ID        int    `gorm:"column:id" json:"id"`
	FName     string `gorm:"column:fname" json:"fname"`
	LName     string `gorm:"column:lname" json:"lname"`
	BirthDate string `gorm:"column:birthdate" json:"birthdate"`
	Age       int    `gorm:"column:age" json:"age"`
	Address   string `gorm:"column:address" json:"address"`
}

func (Test1) TableName() string {
	return "test1"
}
