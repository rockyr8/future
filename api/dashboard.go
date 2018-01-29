package api

import (
	"net/http"

	. "future/model"

	"github.com/gin-gonic/gin"
)


//获取list
func GetDashboardListAPI(c *gin.Context) {
	d := Dashboard{}
	str := d.GetList()
	c.String(http.StatusOK, str)
}



