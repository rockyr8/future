package middleware

import (
	"fmt"
	"net/http"
	"time"

	db "future/database"
	"future/common"

	"github.com/gin-gonic/gin"
)

//屏蔽跨域错误 Access-Control-Allow-Origin  开放post ,get
func OpenMiddleWare(c *gin.Context) {
	// fmt.Println("OpenMiddleWare")
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Next()
}

//token验证 此方法仅验证用户是否登录过期
func AuthMiddleWare(c *gin.Context) {
	uid := c.Request.FormValue("uid")
	ctoken := c.Request.FormValue("token")
	// fmt.Printf("uid is %s, token is %s\n", uid, ctoken)
	if uid != "" && ctoken != "" {
		token := db.RedisGet(uid)
		if token == ctoken {
			//开一个协程进行续租验证
			go renew(uid, token)
			return
		}
	}

	// c.JSON(http.StatusUnauthorized, gin.H{
	// 	"error": "Unauthorized",
	// })

	//终止当前访问
	c.String(http.StatusOK, "Unauthorized")
	c.Abort()
	// fmt.Println("1")
}

//此方法用于鉴权 哪些页面可以访问 哪些数据可以访问
func AuthLogicMiddleWare(c *gin.Context) {
	// fmt.Printf("ip=%v,url=%v",c.ClientIP,c.Request.RequestURI)
	// c.Request.RequestURI 返回  /v2/login
	uid := c.Request.FormValue("uid")
	ctoken := c.Request.FormValue("token")
	// fmt.Printf("uid is %s, token is %s\n", uid, ctoken)
	if uid != "" && ctoken != "" {
		token := db.RedisGet(uid)
		if token == ctoken {
			// fmt.Printf("uid is %s, token is %s\n", uid, ctoken)
			// 这里第二步验证，验证url是否在里面
			if url := db.RedisGet(uid + c.Request.RequestURI); url != "" {
				// fmt.Printf("url is %s,urlval is %s\n", uid+c.Request.RequestURI,url)
				//开一个协程进行续租验证
				go renew(uid, token)
				return
			}

		}
	}

	//终止当前访问
	c.String(http.StatusOK, "Unauthorized")
	c.Abort()
}

//续租方法：剩余5分钟开始续租，续租时间25分钟。
func renew(uid string, token string) {
	ms := int64(db.RedisClient.TTL(uid).Val() / time.Second)
	fmt.Println("renew", ms)
	if ms < int64(common.RemainRenewalTime) && ms > 0 {
		db.RedisSet(uid, token, common.RenewalTime)
		fmt.Println("ms=", ms)
		GetAcceseAuth(uid, common.RenewalTime, 0)
	}
}

//获取 or 删除高级访问权限 用于页面验证和按钮验证 operating 0=>获取 1=>删除
func GetAcceseAuth(uid string, sec time.Duration, operating int) {
	var url string
	rows, err := db.SqlDB.Query(`SELECT DISTINCT authurl FROM go_menu a LEFT JOIN go_menu_role b ON a.id=b.menuID LEFT JOIN go_account c ON b.roleID=c.roleID
		WHERE a.valid=0 AND c.id=? ORDER BY a.sort DESC`, uid)
	defer rows.Close()
	if err != nil {
		//加入错误日志
		return
	}

	var delurl []string
	for rows.Next() {
		rows.Scan(&url)
		//高级权限加入redis，用于接口验证。默认存储20分钟 每15分钟进行一次续租，续租时间25分钟。
		if operating == 0 {
			db.RedisSet(uid+url, "1", sec)
		} else {
			delurl = append(delurl, uid+url) //注销登录时，需要删除的url
		}
	}

	if err = rows.Err(); err != nil {
		//加入错误日志
		return
	}

	//注销时用到
	if operating == 1 {
		db.RedisDel(delurl...)
	}

}
