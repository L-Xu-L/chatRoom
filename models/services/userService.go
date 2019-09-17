package services

import (
	"chatRoom/models/entitys"
	"chatRoom/models/repositories"
	"chatRoom/util"
)

type UserService interface {
	Login(entitys.UserEntity) (*entitys.UserEntity,error)
	GetAllUser() ([]*entitys.UserEntity,error)
	GetUserByOpenid(string) (*entitys.UserEntity,error)
	Save(*entitys.UserEntity) error
}

type userService struct {
	repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (this *userService) Login(user entitys.UserEntity) (*entitys.UserEntity,error) {
	userMap := util.StructToMap(user)
	delete(userMap,"idUser")
	return this.UserRepository.Create(userMap)
}

func (this *userService) GetUserByOpenid(openid string) (*entitys.UserEntity,error) {
	return this.UserRepository.GetUserByOpenid(openid)
}

func (this *userService) GetAllUser() ([]*entitys.UserEntity,error) {
	return this.UserRepository.All()
}

func (this *userService) Save(user *entitys.UserEntity) error {
	 return this.UserRepository.UpdateUserById(user.IdUser, util.StructToMap(*user))
}