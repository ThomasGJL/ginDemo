package controller

import (
	"ginDemo/dto"
	"ginDemo/dao"
	"ginDemo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	userService service.UserService
}

func UserApi(router *gin.Engine) {

	userHandler := UserHandler{
		userService: &service.UserServiceImpl{},
	}

	userGroup := router.Group("user/")
	{
		userGroup.GET("/:id", userHandler.SelectById)
		userGroup.GET("/all", userHandler.SelectAllUsers)
		userGroup.POST("/store", userHandler.StoreUser)
	}
}

// 根据ID查询用户
func (userHandler UserHandler) SelectById(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	user := userHandler.userService.SelectById(userId)

	c.JSON(http.StatusOK, dto.Ok(user))
}

func (userHandler UserHandler) SelectAllUsers(c *gin.Context) {
	var userList []dao.User
	userList = userHandler.userService.SelectAllUsers()

	c.JSON(http.StatusOK, dto.Ok(userList))
}

func (userHandler UserHandler) StoreUser(c *gin.Context) {
	name := c.PostForm("name")
	log.Info(name)
	//user := userHandler.userService.StoreUser()

	c.JSON(http.StatusOK, dto.Ok(name))
}