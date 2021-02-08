package controllers

import (
	"github.com/devmaufh/golang-api-rest/models"
	"github.com/devmaufh/golang-api-rest/services"

	"github.com/gin-gonic/gin"
)

//ModuleController interface
type ModuleController interface {
	FindAll() []models.Module
	Save(ctx *gin.Context) models.Module
}

type moduleController struct {
	service services.ModuleServiceInteface
}

//NewModuleController creates an instance of controller
func NewModuleController(service services.ModuleServiceInteface) ModuleController {
	return &moduleController{
		service: service,
	}
}

func (c *moduleController) FindAll() []models.Module {
	return c.service.FindAll()
}
func (c *moduleController) Save(ctx *gin.Context) models.Module {
	var module models.Module
	ctx.BindJSON(&module)
	c.service.Save(module)
	return module
}
