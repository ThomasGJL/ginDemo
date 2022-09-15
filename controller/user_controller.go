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
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
	}
}

// 根据ID查询用户
func (userHandler UserHandler) SelectById(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	user := userHandler.userService.SelectById(userId)

	c.JSON(http.StatusOK, dto.Ok(user))
}

// 所有用户
func (userHandler UserHandler) SelectAllUsers(c *gin.Context) {
	var userList []dao.User
	userList = userHandler.userService.SelectAllUsers()

	c.JSON(http.StatusOK, dto.Ok(userList))
}

// 新建用户
func (userHandler UserHandler) StoreUser(c *gin.Context) {
	username := c.PostForm("name")
	log.Info("username: ", username)
	nuser := dao.User{Name: username}
	userHandler.userService.StoreUser(nuser)

	c.JSON(http.StatusOK, dto.Ok("success"))
}

// 修改用户
func (userHandler UserHandler) UpdateUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	username := c.PostForm("name")
	//log.Info("userId: ", userId)
	//log.Info("username: ", username)
	cuser := dao.User{Id: userId, Name: username}
	userHandler.userService.UpdateUser(userId, cuser)

	c.JSON(http.StatusOK, dto.Ok("success"))
}

// 删除用户
func (userHandler UserHandler) DeleteUser(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	//log.Info("userId: ", userId)
	duser := dao.User{Id: userId}
	userHandler.userService.DeleteUser(userId, duser)

	c.JSON(http.StatusOK, dto.Ok("success"))
}