//菜单管理
package api

import (
	"net/http"

	. "future/model"

	"github.com/gin-gonic/gin"
)


//获取主菜单list
func GetMainMenuListAPI(c *gin.Context) {
	id := c.PostForm("id")
	num := c.PostForm("num")
	m := Menu{ID:id,ChildNum:num}

	str, err := m.GetList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//添加 or 修改主菜单
func OperateMainMenuAPI(c *gin.Context) {
	mid := c.PostForm("id")
	nickName := c.PostForm("nickname")
	ico := c.PostForm("ico")
	url := c.PostForm("url")
	operate := c.PostForm("operate")
	m := Menu{mid,nickName,ico,url,""}

	var id int64
	var err error
	if operate == "0" {
		id, err = m.Add()
	} else if operate == "1" {
		id, err = m.Modify()
	}

	if err != nil || id < 1 {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

//获取子菜单list
func GetChildMenuListAPI(c *gin.Context) {
	id := c.PostForm("id")
	classid := c.PostForm("classid")
	m := ChildMenu{ID:id,ClassID:classid}
	str, err := m.GetList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//添加 or 修改子菜单
func OperateChildMenuAPI(c *gin.Context) {
	mid := c.PostForm("id")
	classID := c.PostForm("classid")
	nickname := c.PostForm("nickname")
	aurl := c.PostForm("aurl")
	url := c.PostForm("url")
	sort := c.PostForm("sort")
	valid := c.PostForm("valid")
	operate := c.PostForm("operate")
	m := ChildMenu{mid,classID,nickname,aurl,url,sort,valid}

	var id int64
	var err error
	if operate == "0" {
		id, err = m.Add()
	} else if operate == "1" {
		id, err = m.Modify()
	}

	if err != nil || id < 1 {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}