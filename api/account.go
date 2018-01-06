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
	if uname=="" || pwd==""{
		c.String(http.StatusOK, "")
		return
	}
	// str,err := GetAccount(uname,pwd)
	str,err := AccountLogin(uname,pwd)
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
	err := AccountLoginOut(uid,ctoken)
	if err != nil{
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

func GetAccountListAPI(c *gin.Context) {
	str,err := GetAccountList()
	if err != nil{
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//左边菜单导航接口 返回菜单名称和相对路径
func GetMenuAPI(c *gin.Context){
	uid := c.PostForm("uid")
	str,err := GetMenu(uid)
	if err != nil{
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

func GetMenu1API(c *gin.Context){
	uid := "1"
	str,err := GetMenu(uid)
	if err != nil{
		c.String(http.StatusOK, "[]")
		return
	}
	c.String(http.StatusOK, str)
}