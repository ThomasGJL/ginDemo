package service

import (
	"ginDemo/dao"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	SelectById(userId int) *dao.User

	SelectAllUsers() []dao.User

	StoreUser() *dao.User
}

type UserServiceImpl struct {
}

func (UserServiceImpl) SelectById(userId int) *dao.User {
	user := &dao.User{}
	user.SelectById(userId)
	return user
}

func (UserServiceImpl) SelectAllUsers() []dao.User {
	var userList []dao.User
	user := &dao.User{}
	userList = user.SelectAllUsers()
	log.Info("user_service===userList===", userList)
	return userList
}

func (UserServiceImpl) StoreUser() *dao.User {
	user := &dao.User{}
	user.StoreUser()
	return user
}
