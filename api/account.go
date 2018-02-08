//用户管理的接口都在这里
package api

import (
	"net/http"
	// "fmt"

	. "future/model"

	"github.com/gin-gonic/gin"
	"strconv"
)

//登录接口 需要用户名和密码 返回uid和昵称
func AccountLoginAPI(c *gin.Context) {
	uname := c.PostForm("uname")
	pwd := c.PostForm("pwd")
	// fmt.Printf("uname=%s,pwd=%s",uname,pwd)
	if uname == "" || pwd == "" {
		c.String(http.StatusOK, "")
		return
	}
	// str,err := GetAccount(uname,pwd)
	str, err := AccountLogin(uname, pwd)
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//注销接口
func AccountLoginOutAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	ctoken := c.PostForm("token")
	err := AccountLoginOut(uid, ctoken)
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

//获取后台登录列表
func GetAccountListAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	userName := c.PostForm("username")
	nickName := c.PostForm("nickname")
	phone := c.PostForm("phone")
	tel := c.PostForm("tel")
	roleID := c.PostForm("roleID")
	valid := c.PostForm("valid")
	ctime := c.PostForm("ctime")
	ltime := c.PostForm("ltime")
	a := Account{Uid: uid, UserName: userName, NickName: nickName, Phone: phone, Tel: tel, RoleID: roleID, Valid: valid, Createtime: ctime, Logintime: ltime}
	str, err := a.GetList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//添加 or 修改 用户
func OperateAccountAPI(c *gin.Context) {
	selfID, err := strconv.Atoi(c.PostForm("uid"))
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	uid := c.PostForm("userid")
	userName := c.PostForm("username")
	passWD := c.PostForm("pwd")
	nickName := c.PostForm("nickname")
	phone := c.PostForm("phone")
	tel := c.PostForm("tel")
	roleID := c.PostForm("roleID")
	valid := c.PostForm("valid")
	operate := c.PostForm("operate")
	proportions, err := strconv.ParseFloat(c.PostForm("proportions"), 64)
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	a := Account{ID: selfID, Uid: uid, UserName: userName, PassWD: passWD, NickName: nickName, Phone: phone, Tel: tel, RoleID: roleID, Valid: valid, Proportions: proportions}

	var id int64
	//var err error
	if operate == "0" {
		id, err = a.Add()
	} else if operate == "1" {
		id, err = a.Modify()
	}

	if err != nil || id < 1 {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")

}

//获取单个用户详情
func GetAccountDetailAPI(c *gin.Context) {
	uid := c.PostForm("userid")
	a := Account{Uid: uid}
	str, err := a.GetDetail()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//获取单个用户银行详情
func GetAccountBankDetailAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	a := AccountBank{AccountID: uid}
	str, err := a.GetBankDetail()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//修改用户银行信息
func OperateAccountBankAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	accountName := c.PostForm("name")
	accountCardNum := c.PostForm("card")
	openBank := c.PostForm("openbank")
	branchBank := c.PostForm("branchbank")
	nickName := c.PostForm("nickname")
	phone := c.PostForm("phone")
	tel := c.PostForm("tel")
	a := Account{NickName:nickName,Phone:phone,Tel:tel}
	bank := AccountBank{AccountID: uid, AccountName: accountName, AccountCardNum: accountCardNum, OpenBank: openBank, BranchBank: branchBank,Account:a}
	rows, err := bank.UpdateBank()
	if err != nil || rows < 1 {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

//修改密码
func ModifyPwdAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	oldpwd := c.PostForm("oldpwd")
	newpwd := c.PostForm("newpwd")
	a := Account{Uid: uid, PassWD: newpwd}
	str := a.ModifyPwd(oldpwd)
	if str == "0" {
		c.String(http.StatusOK, "SUCCESS")
	}
}

//左边菜单导航接口 返回菜单名称和相对路径
func GetMenuAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	a := Account{Uid: uid}
	str, err := a.GetMenu()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//生成等级关联 用于数据权限 自己可以看到自己和后代的数据
func CreateChildAPI(c *gin.Context) {
	err := CreateChild()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

//分成表 生成昨天的分成金额
func CreateSettlementAPI(c *gin.Context) {
	err := CreateSettlement()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}
