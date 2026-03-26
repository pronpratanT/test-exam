package models

type Test2 struct {
	ID        int    `gorm:"column:id" json:"id"`
	Username  string `gorm:"column:username" json:"username"`
	Password  string `gorm:"column:password" json:"password"` // เก็บ hash
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
}

func (Test2) TableName() string {
	return "test2"
}
