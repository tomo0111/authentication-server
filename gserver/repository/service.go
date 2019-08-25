package repository

import (
	"github.com/tomoyane/grant-n-z/gserver/entity"
	"github.com/tomoyane/grant-n-z/gserver/model"
)

type ServiceRepository interface {
	FindAll() ([]*entity.Service, *model.ErrorResponse)

	FindById(id int) (*entity.Service, *model.ErrorResponse)

	FindByName(name string) (*entity.Service, *model.ErrorResponse)

	FindByApiKey(apiKey string) (*entity.Service, *model.ErrorResponse)

	Save(service entity.Service) (*entity.Service, *model.ErrorResponse)

	Update(service entity.Service) *entity.Service
}

