package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/widuu/gojson"
)

//设置websocket
//CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//websocket实现
func ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据，message为byte
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		send_message := ""
		json := string(message[:])
		retcode := gojson.Json(json).Get("retcode").Tostring()

		if retcode != "404" {
			//fmt.Println(json)
			message_type := gojson.Json(json).Get("message_type").Tostring()
			if message_type == "guild" {
				guild_id := gojson.Json(json).Get("guild_id").Tostring()
				channel_id := gojson.Json(json).Get("channel_id").Tostring()
				message_text := gojson.Json(json).Get("message").Tostring()
				self_id := gojson.Json(json).Get("self_id").Tostring()
				self_tiny_id := gojson.Json(json).Get("self_tiny_id").Tostring()
				user_id := gojson.Json(json).Get("user_id").Tostring()
				time := gojson.Json(json).Get("time").Tostring()
				fmt.Println("[Time-" + time + "]收到来自频道分组：" + guild_id + "下子频道：" + channel_id + "内用户：" + user_id + "发送的消息：" + message_text)
				fmt.Println("来源机器人：" + self_tiny_id + "[" + self_id + "]")

				if message_text == "hello" {
					//收到来自频道 Harmonica(40759201636795692) 子频道 测试频道(1392839) 内 Harmonica(144115218677824752) 的消息: 1
					send_message = `{"action":"send_guild_channel_msg","params":{"guild_id":40759201636795692,"channel_id":1392839,"message":"hello world"}}`
				}
				fmt.Println(send_message)
			}
		}

		returnbyte := []byte(send_message)
		err = ws.WriteMessage(mt, returnbyte)
		if err != nil {
			break
		}

	}
}

func main() {
	r := gin.Default()
	r.GET("/ws", ping)
	r.Run(":7790")
}
