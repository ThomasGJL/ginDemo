package controller

import (
	"ginDemo/dto"
	"ginDemo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
		userGroup.GET("/:id", userHandler.user)
	}
}

// 根据ID查询用户
func (userHandler UserHandler) user(c *gin.Context) {
	userIdStr := c.Param("id")
	userId, _ := strconv.Atoi(userIdStr)
	user := userHandler.userService.User(userId)

	c.JSON(http.StatusOK, dto.Ok(user))
}
