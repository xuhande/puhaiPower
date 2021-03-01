package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"puhai/server/service"
)

func ViewData(c *gin.Context) {
	powerLoad := service.GetData()
	c.IndentedJSON(http.StatusOK, gin.H{
		"title": gin.H {
			"text": "电力负荷图",
		},
		"tooltip": "",
		"xAxis": gin.H{
			"data": []string{"衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"},
			"name": "产品",
		},
		"series": gin.H {
			"name": "销量",
			"data": powerLoad,
			"itemStyle": gin.H {
				"normal": gin.H {
					"color": "hotpink",
				},
			},
		},
	})
}