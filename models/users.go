package models

type User struct {
	Email string `json:"email"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
    Token string `json:"log"`
}

type Response struct {
	Error string `json:"error"`
	Result string `json:"result"`
}