package models

//LoginCredentials model for manage login data
type LoginCredentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
