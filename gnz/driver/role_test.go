package driver

import (
	"testing"

	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/tomoyane/grant-n-z/gnz/entity"
	"github.com/tomoyane/grant-n-z/gnz/log"
)

var roleRepository RoleRepository

// Setup test precondition
func init() {
	log.InitLogger("info")

	db, _ := gorm.Open("sqlite3", "/tmp/test_grant_nz.db")
	connection = db
	roleRepository = NewRoleRepository()
}

// FindAll InternalServerError test
func TestRoleFindAll_InternalServerError(t *testing.T) {
	_, err := roleRepository.FindAll()
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleFindAll_InternalServerError test")
	}
}

// FindOffSetAndLimit InternalServerError test
func TestRoleFindOffSetAndLimit_InternalServerError(t *testing.T) {
	_, err := roleRepository.FindOffSetAndLimit(1, 1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleFindOffSetAndLimit_InternalServerError test")
	}
}

// FindById InternalServerError test
func TestRoleFindById_InternalServerError(t *testing.T) {
	_, err := roleRepository.FindById(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleFindById_InternalServerError test")
	}
}

// FindByName InternalServerError test
func TestRoleFindByName_InternalServerError(t *testing.T) {
	_, err := roleRepository.FindByName("test")
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleFindByName_InternalServerError test")
	}
}

// FindByNames InternalServerError test
func TestRoleFindByNames_InternalServerError(t *testing.T) {
	_, err := roleRepository.FindByNames([]string{"test"})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleFindByNames_InternalServerError test")
	}
}

// FindByGroupId InternalServerError test
func TestRoleFindByGroupId_InternalServerError(t *testing.T) {
	_, err := roleRepository.FindByGroupId(1)
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleFindByGroupId_InternalServerError test")
	}
}

// FindNameById is nil test
func TestRoleFindNameById_Nil(t *testing.T) {
	name := roleRepository.FindNameById(1)
	if name != nil {
		t.Errorf("Incorrect TestRoleFindNameById_Nil test")
	}
}

// Save InternalServerError test
func TestRoleSave_InternalServerError(t *testing.T) {
	_, err := roleRepository.Save(entity.Role{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleSave_InternalServerError test")
	}
}

// SaveWithRelationalData InternalServerError test
func TestRoleSaveWithRelationalData_InternalServerError(t *testing.T) {
	_, err := roleRepository.SaveWithRelationalData(1, entity.Role{})
	if err.Code != http.StatusInternalServerError {
		t.Errorf("Incorrect TestRoleSaveWithRelationalData_InternalServerError test")
	}
}
