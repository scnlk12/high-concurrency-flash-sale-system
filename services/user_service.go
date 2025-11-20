package services

import (
	"github.com/scnlk12/high-concurrency-flash-sale-system/datamodels"
	"github.com/scnlk12/high-concurrency-flash-sale-system/repositories"
	"golang.org/x/crypto/bcrypt"
)

// 定义接口
type IUserService interface {
	// 方法
	IsPwdSuccess(userName, pwd string) (*datamodels.User, bool)
	AddUser(user *datamodels.User) (int64, error)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func NewUserService(userRepository repositories.IUserRepository) IUserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (u *UserService) IsPwdSuccess(userName, pwd string) (*datamodels.User, bool) {
	user, err := u.UserRepository.Select(userName)
	if err != nil {
		return nil, false
	}
	
	isOk, _ := ValidPassword(pwd, user.HashPassword)
	if !isOk {
		return nil, false
	}
	return user, isOk
}

// 判断密码是否匹配
func ValidPassword(userPassword, hashed string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func (u *UserService) AddUser(user *datamodels.User) (int64, error) {
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return 0, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.Insert(user)
}

// 密码加密
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}