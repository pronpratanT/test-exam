package dto

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

type SigninResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Test4DTO struct {
	FName      string `json:"fname"`
	LName      string `json:"lname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Profile    string `json:"profile"`
	Birthdate  string `json:"birthdate"`
	Occupation string `json:"occupation"`
}
