package service

import (
	"github.com/satori/go.uuid"
	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
)

type RoleService struct {
	RoleRepository repository.RoleRepository
}

func NewRoleService() RoleService {
	return RoleService{RoleRepository: repository.RoleRepositoryImpl{}}
}

func (rs RoleService) InsertRole(role *entity.Role) (*entity.Role, *entity.ErrorResponse) {
	role.Uuid, _ = uuid.NewV4()
	return rs.RoleRepository.Save(*role)
}