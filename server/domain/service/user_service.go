package service

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/tomoyane/grant-n-z/server/domain/entity"
	"github.com/tomoyane/grant-n-z/server/domain/repository"
	"github.com/tomoyane/grant-n-z/server/log"
)

var logger log.Log

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUserService() UserService {
	logger = log.NewLogger()
	return UserService{UserRepository: repository.UserRepositoryImpl{}}
}

func (us UserService) EncryptPw(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([] byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func (us UserService) ComparePw(passwordHash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		logger.Error("error compare password", err)
		return false
	}

	return true
}

func (us UserService) InsertUser(user entity.User) (*entity.User, *entity.ErrorResponse) {
	user.Uuid, _ = uuid.NewV4()
	user.Password = us.EncryptPw(user.Password)
	return us.UserRepository.Save(user)
}