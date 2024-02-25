package route

import (
	"github.com/gin-gonic/gin"
	"photo_album/service"
)

func InitUserRoute(r *gin.Engine) {
	// 新增用户
	r.PUT("/user", service.AddUser)
	// 删除用户
	r.DELETE("/user/:id", service.DeleteUser)
	// TODO 修改用户信息

	// TODO 查询用户信息
}
