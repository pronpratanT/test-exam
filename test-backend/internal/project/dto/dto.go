package dto

type Test1DTO struct {
	FName     string `json:"fname"`
	LName     string `json:"lname"`
	BirthDate string `json:"birthdate"`
	Age       int    `json:"age"`
	Address   string `json:"address"`
}
