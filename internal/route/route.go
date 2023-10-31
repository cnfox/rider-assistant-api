package route

import (
	"snail/internal/handler/event"
	"snail/internal/handler/rider"
	"snail/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() *gin.Engine {
	app := gin.New()
	// 中间件
	app.Use(gin.Logger(), middleware.Cors(), middleware.RecoverAtLast(), middleware.TraceId())
	// 接口不存在
	app.NoRoute(middleware.NotFound())
	// 路由分组
	api := app.Group("/api/")

	// Welcome
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Welcome": "This is rider assistant api",
		})
	})

	// 队员
	api.POST("/rider/login", rider.Login)   // 队员登录
	api.GET("/rider/info", rider.Info)      // 队员详情
	api.POST("/rider/update", rider.Update) // 更改个人信息

	// 活动
	api.GET("/event/list", event.List)        // 车队活动列表
	api.GET("/event/info", event.Info)        // 车队活动详情
	api.POST("/event/enroll", event.Enroll)   // 队员活动报名
	api.POST("/event/quit", event.Quit)       // 队员退出活动
	api.POST("/event/checkin", event.Checkin) // 队员活动签到

	return app
}
