package main

import (
	. "future/api"
	. "future/logic"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境gin.DebugMode，线上环境为gin.ReleaseMode
	router := gin.Default()

	//注册路由 person 例子
	router.GET("/", IndexApi)
	router.GET("/persons/:key", GetRedisValAPI)
	router.GET("/set", SetRedisValAPI)
	// router.GET("/login/:key", AccountLoginApi)
	
	//添加全局中间件(在这行代码之前设置的路由规则,不经过g该中间件) => 屏蔽跨域错误 Access-Control-Allow-Origin
	router.Use(OpenMiddleWare)
	//注册路由
	router.POST("/login", AccountLoginAPI)
	router.POST("/logout", AccountLoginOutAPI)

	//添加分组中间件=>对用户进行简单认证，通常是权限不太严格的接口或者是公共接口
	//v1内容都是需要登录才能访问
	authorized := router.Group("/v1")
	authorized.Use(AuthMiddleWare)
	{
		// authorized.POST("/logout", AccountLoginOutAPI)
		authorized.POST("/menu", GetMenuAPI)
	}

	//添加分组中间件=>对用户进行鉴权 相对严格的权限接口
	//v2内容除了登录，还需要判断用户有没有接口访问的权限
	authorized = router.Group("/v2")
	authorized.Use(AuthLogicMiddleWare)
	{
		authorized.POST("/account", GetAccountListAPI)
	}
	
	// router.GET("/person/:id", GetPersonApi)
	// router.PUT("/person/:id", ModPersonApi)
	// router.DELETE("/person/:id", DelPersonApi)

	return router
}
