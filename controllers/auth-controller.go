package controllers

import (
	"github.com/devmaufh/golang-api-rest/models"
	"github.com/devmaufh/golang-api-rest/services"
	"github.com/gin-gonic/gin"
)

//LoginController Defines an interface to consume login services
type LoginController interface {
	Login(ctx *gin.Context) (int, map[string]string)
}

type loginController struct {
	loginService services.LoginServiceI
	jwtService   services.JWTServiceI
}

//NewLoginController creates new Login controller
func NewLoginController(loginService services.LoginServiceI,
	jWtService services.JWTServiceI) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) (int, map[string]string) {
	var credentials models.LoginCredentials
	err := ctx.ShouldBindJSON(&credentials)
	if err != nil {
		return 400, map[string]string{"error": "bad request"}
	}
	isUserAuthenticated := controller.loginService.LoginUser(credentials.Email, credentials.Password)
	if !isUserAuthenticated {
		return 401, map[string]string{"error": "Invalid credentials"}
	}
	return 200, map[string]string{"access_token": controller.jwtService.GenerateToken(credentials.Email, true)}

}
