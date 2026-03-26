package models

type Test4 struct {
	ID         int    `gorm:"column:id" json:"id"`
	FName      string `gorm:"column:fname" json:"fname"`
	LName      string `gorm:"column:lname" json:"lname"`
	Email      string `gorm:"column:email" json:"email"`
	Phone      string `gorm:"column:phone" json:"phone"`
	Profile    string `gorm:"column:profile" json:"profile"`
	Birthdate  string `gorm:"column:birthdate" json:"birthdate"`
	Occupation string `gorm:"column:occupation" json:"occupation"`
	Sex        string `gorm:"column:sex" json:"sex"`
	CreatedAt  string `gorm:"column:created_at" json:"created_at"`
}

func (Test4) TableName() string {
	return "test4"
}
