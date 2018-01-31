//用户管理的接口都在这里
package api

import (
	"net/http"
	// "fmt"

	. "future/model"

	"github.com/gin-gonic/gin"
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
	a := Account{}
	str, err := a.GetList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//添加 or 修改 用户
func OperateAccountAPI(c *gin.Context) {
	uid := c.PostForm("userid")
	userName := c.PostForm("username")
	passWD := c.PostForm("pwd")
	nickName := c.PostForm("nickname")
	phone := c.PostForm("phone")
	tel := c.PostForm("tel")
	roleID := c.PostForm("roleID")
	valid := c.PostForm("valid")
	operate := c.PostForm("operate")
	a := Account{Uid:uid, UserName:userName, PassWD:passWD, NickName:nickName, Phone:phone, Tel:tel, RoleID:roleID, Valid:valid}

	var id int64
	var err error
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

//修改密码
func ModifyPwdAPI(c *gin.Context){
	uid := c.PostForm("uid")
	oldpwd := c.PostForm("oldpwd")
	newpwd := c.PostForm("newpwd")
	a := Account{Uid:uid,PassWD:newpwd}
	str := a.ModifyPwd(oldpwd)
	if str == "0" {
		c.String(http.StatusOK, "SUCCESS")
	}
}


//左边菜单导航接口 返回菜单名称和相对路径
func GetMenuAPI(c *gin.Context) {
	uid := c.PostForm("uid")
	a := Account{Uid:uid}
	str, err := a.GetMenu()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//左边菜单导航接口 返回菜单名称和相对路径
func CreateChildAPI(c *gin.Context) {
	err := CreateChild()
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}




