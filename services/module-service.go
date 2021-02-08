package services

import (
	"github.com/devmaufh/golang-api-rest/models"
)

//ModuleServiceInteface build and interface for access to service implementation
type ModuleServiceInteface interface {
	Save(models.Module) models.Module
	FindAll() []models.Module
}

type moduleService struct {
	modules []models.Module
}

//New creates new ModuleServiceInteface
func New() ModuleServiceInteface {
	return &moduleService{
		modules: []models.Module{},
	}
}

func (service *moduleService) Save(module models.Module) models.Module {
	service.modules = append(service.modules, module)
	return module
}

func (service *moduleService) FindAll() []models.Module {
	return service.modules
}
