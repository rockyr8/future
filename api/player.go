package api

import (
	"net/http"

	. "future/model"
	. "future/server"

	"github.com/gin-gonic/gin"
	"strconv"
)


//获取list
func GetPlayerListAPI(c *gin.Context) {
	id := c.PostForm("id")
	p := Player{ID:id}
	str, err := p.GetList()
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, str)
}

//启用 禁用
func DisablePlayerAPI(c *gin.Context) {
	playerid := c.PostForm("id")
	valid := c.PostForm("valid")
	p := Player{ID:playerid,Valid:valid}

	rows, err := p.Disable()

	if err != nil || rows < 1 {
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}

//玩家上分
func AddMoney(c *gin.Context) {
	id := c.PostForm("id")
	points := c.PostForm("points")
	playerid,_ := strconv.Atoi(id)
	gold,_ := strconv.Atoi(points)
	g := GameServer{PlayerID:playerid,Gold:gold}
	if(!g.AddMoney()){
		c.String(http.StatusOK, "")
		return
	}
	c.String(http.StatusOK, "SUCCESS")
}