package dao

import (
	"ginDemo/config/db"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Id   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"size:50"`
}

func (user *User) SelectById(userId int) {
	db.DB.First(&user, userId)
}

func (user *User) SelectAllUsers() []User {
	var userList []User
	db.DB.Debug().Find(&userList)
	log.Info("user_dao===userList===", userList)
	return userList
}

func (user *User) StoreUser() {
	db.DB.Create(&user)
}