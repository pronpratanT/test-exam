package dto

import "time"

type Test1DTO struct {
	FName     string `json:"fname"`
	LName     string `json:"lname"`
	BirthDate string `json:"birthdate"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}

type Test2DTOSignup struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Test2DTOSignin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Test3DTO struct {	
	Name   string `json:"name"`
	Reason string `json:"reason"`
	Status string `json:"status"`
}

type Test3ApproveDTO struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
}

type Test3RejectDTO struct {
	ID     int    `json:"id"`
	Reason string `json:"reason"`
}

type SigninResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Test4DTO struct {
	FName      string    `form:"fname" json:"fname"`
	LName      string    `form:"lname" json:"lname"`
	Email      string    `form:"email" json:"email"`
	Phone      string    `form:"phone" json:"phone"`
	Profile    string    `form:"-" json:"profile"`
	Birthdate  time.Time `form:"birthdate" time_format:"2006-01-02" json:"birthdate"`
	Occupation string    `form:"occupation" json:"occupation"`
	Sex        string    `form:"sex" json:"sex"`
}

type Test5DTO struct {
	QueNumber string `json:"que_number"`
	CreatedAt string `json:"created_at"`
	IsCleared bool   `json:"iscleared"`
}