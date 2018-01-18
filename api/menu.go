//菜单、角色管理
package api

import (
	"net/http"

	. "future/model"

	"github.com/gin-gonic/gin"
)


//获取主菜单list
func GetMainMenuListAPI(c *gin.Context) {
	id := c.PostForm("id")
	m := Menu{}
	str, err := m.GetList(id)
	if err != nil {
		c.String(http.StatusOK, "111")
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
	m := Menu{mid,nickName,ico,url}

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