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
	//log.Info("user_dao===userList===", userList)
	return userList
}

func (user *User) StoreUser(nuser User) {
	log.Info("user_dao===StoreUser===", nuser.Name)
	newUser := User{Name:nuser.Name}
	db.DB.Debug().Create(&newUser)
}

func (user *User) UpdateUser(cuser User) {
	db.DB.Debug().Save(cuser)
}

func (user *User) DeleteUser(duser User) {
	db.DB.Debug().Delete(duser)
}