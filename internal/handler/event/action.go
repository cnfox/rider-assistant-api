package event

import (
	"snail/pkg/response"

	"github.com/gin-gonic/gin"
)

// 车队活动列表
func List(c *gin.Context) {
	var data interface{} = nil
	response.Succuss(c, data)
}

// 车队活动详情
func Info(c *gin.Context) {
	var data interface{} = nil
	response.Succuss(c, data)
}

// 队员活动报名
func Enroll(c *gin.Context) {
	// 活动前一日凌晨停止报名
	var data interface{} = nil
	response.Succuss(c, data)
}

// 队员退出活动
func Quit(c *gin.Context) {
	// 活动前一日凌晨后不能退出
	var data interface{} = nil
	response.Succuss(c, data)
}

// 队员活动签到
func Checkin(c *gin.Context) {
	// 活动当日签到，发放基础积分，记录流水
	var data interface{} = nil
	response.Succuss(c, data)
}
