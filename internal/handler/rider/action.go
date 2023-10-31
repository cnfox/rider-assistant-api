package rider

import (
	"snail/pkg/response"

	"github.com/gin-gonic/gin"
)

// 队员登录
func Login(c *gin.Context) {
	// 接受 code，换取 openid
	// 获取或创建新用户
	// 返回用户信息
	var data interface{} = nil
	response.Succuss(c, data)
}

// 队员详情
func Info(c *gin.Context) {
	var data interface{} = nil
	response.Succuss(c, data)
}

// 更改个人信息
func Update(c *gin.Context) {
	// 更新花名和分队
	var data interface{} = nil
	response.Succuss(c, data)
}
