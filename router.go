package main

import (
	. "future/api"
	. "future/middleware"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境gin.DebugMode，线上环境为gin.ReleaseMode
	router := gin.Default()

	//注册路由 person 例子
	router.GET("/", IndexApi)
	router.GET("/redisV/:key", GetRedisValAPI)
	router.GET("/redisT/:key", GetRedisValTimeAPI)
	router.GET("/set", SetRedisValAPI)

	//test
	router.GET("/cd", CreateChildAPI)

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
		/* region 系统用户管理*/
		//查询用户
		authorized.POST("/account", GetAccountListAPI)
		//查询用户
		authorized.POST("/account/detail", GetAccountDetailAPI)
		//添加用户
		authorized.POST("/account/add", OperateAccountAPI)
		//修改用户
		authorized.POST("/account/modify", OperateAccountAPI)
		//修改密码
		authorized.POST("/account/modifypwd", ModifyPwdAPI)

		//生成等级关系
		authorized.POST("/account/createchild", CreateChildAPI)

		/* endregion 系统用户管理*/


		/* region 系统权限管理*/

		//主菜单列表
		authorized.POST("/menu/menulist", GetMainMenuListAPI)
		//主菜单新增
		authorized.POST("/menu/add", OperateMainMenuAPI)
		//主菜单修改
		authorized.POST("/menu/modify", OperateMainMenuAPI)

		//子菜单列表
		authorized.POST("/menu/menulistc", GetChildMenuListAPI)
		//子菜单新增
		authorized.POST("/menu/addc", OperateChildMenuAPI)
		//子菜单修改
		authorized.POST("/menu/modifyc", OperateChildMenuAPI)

		//角色列表
		authorized.POST("/role", GetRoleListAPI)
		//角色添加
		authorized.POST("/role/add", OperateRoleAPI)
		//角色修改
		authorized.POST("/role/modify", OperateRoleAPI)
		//角色菜单列表
		authorized.POST("/role/menu", GetRoleMenuListAPI)
		//修改角色菜单
		authorized.POST("/role/menu/modify", ModifyRoleMenuAPI)

		/* endregion 系统用户管理*/

		/* region 玩家管理*/

		//玩家列表
		authorized.POST("/player", GetPlayerListAPI)
		//玩家禁用，启用
		authorized.POST("/player/disable", DisablePlayerAPI)
		//玩家上分操作
		authorized.POST("/player/addpoints", AddMoney)

		/* endregion 玩家管理*/


		/* region 統計*/

		//authorized.POST("/dashboard", GetDashboardListAPI)
		authorized.POST("/dashboard/newplayer", GetDashboardNewAPI)
		authorized.POST("/dashboard/loginplayer", GetDashboardLoginAPI)
		authorized.POST("/dashboard/logincount", GetDashboardLogincountAPI)
		authorized.POST("/dashboard/onlineplayer", GetDashboardOnlineAPI)
		authorized.POST("/dashboard/signplayer", GetDashboardSignAPI)
		authorized.POST("/dashboard/coreplayer", GetDashboardCoreAPI)
		authorized.POST("/dashboard/onlinetime", GetDashboardOnlineTimeAPI)
		authorized.POST("/dashboard/crecharge", GetDashboardCRechargeAPI)
		authorized.POST("/dashboard/trecharge", GetDashboardTRechargeAPI)
		authorized.POST("/dashboard/jackport", GetDashboardJackpotAPI)
		authorized.POST("/dashboard/balance", GetDashboardBalanceAPI)

		/* endregion 統計*/

	}

	// router.GET("/person/:id", GetPersonApi)
	// router.PUT("/person/:id", ModPersonApi)
	// router.DELETE("/person/:id", DelPersonApi)

	return router
}