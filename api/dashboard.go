package api

import (
	"net/http"

	. "future/model"

	"github.com/gin-gonic/gin"
	"time"
)

////获取list
//func GetDashboardListAPI(c *gin.Context) {
//	d := Dashboard{}
//	str := d.GetList()
//	c.String(http.StatusOK, str)
//}

//获取新用户
func GetDashboardNewAPI(c *gin.Context) {
	t := time.Now().Format("20060102")
	str := GetNum("SELECT COUNT(0) FROM kbe_accountinfos WHERE FROM_UNIXTIME(regtime,'%Y%m%d')='" + t + "'")
	c.String(http.StatusOK, str)
}

//获取登录人数
func GetDashboardLoginAPI(c *gin.Context) {
	t := time.Now().Format("20060102")
	str := GetNum("SELECT COUNT(DISTINCT sm_avatarID) FROM tbl_SysLogs	WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='" + t + "'")
	c.String(http.StatusOK, str)
}

//获取登录次数
func GetDashboardLogincountAPI(c *gin.Context) {
	t := time.Now().Format("20060102")
	str := GetNum("SELECT COUNT(sm_avatarID) FROM tbl_SysLogs WHERE sm_logType=10 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='" + t + "'")
	c.String(http.StatusOK, str)
}

//获取在线人数
func GetDashboardOnlineAPI(c *gin.Context) {
	t := time.Now().Format("20060102")
	str := GetNum("SELECT count(0) FROM tbl_SysLogs WHERE sm_logType=89 AND FROM_UNIXTIME(sm_date,'%Y%m%d')='" + t + "'")
	c.String(http.StatusOK, str)
}

//获取签到人数
func GetDashboardSignAPI(c *gin.Context) {
	t := time.Now().Format("20060102")
	str := GetNum("SELECT COUNT(DISTINCT sm_avatarid) FROM tbl_Sign WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')='" + t + "'")
	c.String(http.StatusOK, str)
}

//获取核心会员
func GetDashboardCoreAPI(c *gin.Context) {
	str := GetNum(`SELECT count(0) FROM tbl_Avatar a LEFT JOIN tbl_Account_characters_values b ON a.sm_dbid=b.sm_dbid LEFT JOIN kbe_accountinfos c ON b.parentID=c.entityDBID
								WHERE c.regtime>1466713420 AND a.id NOT IN (112,292,333,324,254,253)
								AND c.lasttime BETWEEN (UNIX_TIMESTAMP(NOW())-604800)  AND UNIX_TIMESTAMP(NOW())
								AND c.numlogin>7
								ORDER BY c.lasttime DESC`)
	c.String(http.StatusOK, str)
}

//获取玩家时长
func GetDashboardOnlineTimeAPI(c *gin.Context) {
	d := Dashboard{}
	str, err := d.GetOnlineTimeLen()
	if err != nil {
		c.String(http.StatusOK, "")
	}
	c.String(http.StatusOK, str)
}

//获取当天充值金額
func GetDashboardCRechargeAPI(c *gin.Context) {
	t := time.Now().Format("20060102")
	str := GetNum("SELECT IFNULL(SUM(sm_price),0) FROM tbl_OnlinePay WHERE FROM_UNIXTIME(sm_date,'%Y%m%d')='"+t+"'")
	c.String(http.StatusOK, str)
}

//获取充值總額
func GetDashboardTRechargeAPI(c *gin.Context) {
	str := GetNum("SELECT IFNULL(SUM(sm_price),0) FROM tbl_OnlinePay")
	c.String(http.StatusOK, str)
}

//获取奖池
func GetDashboardJackpotAPI(c *gin.Context) {
	d := Dashboard{}
	str, err := d.GetJackpot()
	if err != nil {
		c.String(http.StatusOK, "")
	}
	c.String(http.StatusOK, str)
}

//获取 公司收入，支出
func GetDashboardBalanceAPI(c *gin.Context) {
	d := Dashboard{}
	str, err := d.GetGoldAndDiamond()
	if err != nil {
		c.String(http.StatusOK, "")
	}
	c.String(http.StatusOK, str)
}