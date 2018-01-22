package api

import (
	"net/http"

	. "future/model"

	"github.com/gin-gonic/gin"
)


//获取list
func GetRoleListAPI(c *gin.Context) {
	id := c.PostForm("id")
	r := Role{ID:id}
	str, err := r.GetList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//添加 or 修改
func OperateRoleAPI(c *gin.Context) {
	mid := c.PostForm("id")
	nickName := c.PostForm("nickname")
	valid := c.PostForm("valid")
	operate := c.PostForm("operate")
	r := Role{ID:mid,NickName:nickName,Valid:valid}

	var id int64
	var err error
	if operate == "0" {
		id, err = r.Add()
	} else if operate == "1" {
		id, err = r.Modify()
	}

	if err != nil || id < 1 {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

//获取list
func GetRoleMenuListAPI(c *gin.Context) {
	id := c.PostForm("id")
	r := Role{ID:id}
	str, err := r.GetRoleMenuList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//修改角色菜单
func ModifyRoleMenuAPI(c *gin.Context) {
	id := c.PostForm("id")
	updateIds := c.PostForm("ids")
	r := Role{ID:id,UpdateIDs:updateIds}
	err := r.ModifyRoleMenu()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

