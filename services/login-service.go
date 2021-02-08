package services

//LoginServiceI creates an interface
type LoginServiceI interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

//LoginService creates an instance of LoginService
func LoginService() LoginServiceI {
	return &loginInformation{
		email:    "user",
		password: "123",
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}
