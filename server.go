package main

import (
	"io"
	"os"

	"github.com/devmaufh/golang-api-rest/controllers"
	"github.com/devmaufh/golang-api-rest/middleware"
	"github.com/devmaufh/golang-api-rest/services"
	"github.com/gin-gonic/gin"
)

var (
	loginService  services.LoginServiceI         = services.LoginService()
	jwtService    services.JWTServiceI           = services.JWTAuthService()
	moduleService services.ModuleServiceInteface = services.New()

	moduleController controllers.ModuleController = controllers.NewModuleController(moduleService)
	loginController  controllers.LoginController  = controllers.NewLoginController(loginService, jwtService)
)

func setupLogOutuput() {
	f, _ := os.Create("logs/turnomatik.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()
	setupLogOutuput()
	server.Use(gin.Recovery(), middleware.Logger())

	//Public routes
	authRoutes := server.Group("/auth")
	authRoutes.POST("login", func(ctx *gin.Context) {
		ctx.JSON(loginController.Login(ctx))
	})

	//Private routes

	moduleRoutes := server.Group("/")
	moduleRoutes.Use(middleware.AuthorizeJWT())

	moduleRoutes.GET("/module", func(ctx *gin.Context) {
		ctx.JSON(200, moduleController.FindAll())
	})
	moduleRoutes.POST("/module", func(ctx *gin.Context) {
		ctx.JSON(201, moduleController.Save(ctx))
	})
	server.Run(":8000")
}
