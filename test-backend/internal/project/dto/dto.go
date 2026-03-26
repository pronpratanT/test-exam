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

type SigninResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
