package service

import (
	"ginDemo/dao"
	log "github.com/sirupsen/logrus"
)

type UserService interface {
	SelectById(userId int) *dao.User

	SelectAllUsers() []dao.User

	StoreUser(nuser dao.User) *dao.User

	UpdateUser(userId int, cuser dao.User) *dao.User

	DeleteUser(userId int, duser dao.User) *dao.User
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

func (UserServiceImpl) StoreUser(nuser dao.User) *dao.User {
	user := &dao.User{}
	user.StoreUser(nuser)
	return user
}

func (UserServiceImpl) UpdateUser(userId int, cuser dao.User) *dao.User {
	user := &dao.User{}
	user.SelectById(userId)
	if user.Id == 0 {
		user = nil
		log.Error("user not found")
	} else {
		user.UpdateUser(cuser)
	}

	return user
}

func (UserServiceImpl) DeleteUser(userId int, duser dao.User) *dao.User {

	user := &dao.User{}
	user.SelectById(userId)
	if user.Id == 0 {
		user = nil
		log.Error("user not found")
	} else {
		user.DeleteUser(duser)
	}

	return nil
}
