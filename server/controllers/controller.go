package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"puhai/server/service"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
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

func WsHandler(c *gin.Context) {
	var (
		wsConn *websocket.Conn
		err error
		data []byte
		pl []float32
	)



	if wsConn, err = upgrader.Upgrade(c.Writer, c.Request, nil);err != nil {
		return // 获取连接失败直接返回
	}

	for {
		pl = service.GetData()
		data, _ = json.Marshal(pl)
		if err = wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR // 发送消息失败，关闭连接
		}
	}

ERR:
	// 关闭连接
	wsConn.Close()

}