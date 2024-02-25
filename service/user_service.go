package service

import (
	"github.com/gin-gonic/gin"
	"photo_album/config"
	"photo_album/dao"
	"photo_album/model"
	"photo_album/utool"
	"strconv"
	"time"
)

func AddUser(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		JSONStatusBadRequest(c, err)
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	// 加密密码
	user.Password = utool.Md5Crypt(user.Password, config.SaltPassword)
	err = dao.AddUser(&user)
	if err != nil {
		JSONStatusInternalServerError(c, err, "数据库操作异常")
		return
	}
	JSONStatusOk(c, "新增用户成功", nil)
	// TODO 给新用户的邮箱发送注册成功的邮件
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	id, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		JSONStatusBadRequest(c, err)
		return
	}
	dao.DeleteUser(id)

}
