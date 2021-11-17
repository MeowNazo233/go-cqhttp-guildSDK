package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//创建消息格式
type guild_message_sender struct {
	Nickname string `json:"nickname"`
	User_id  uint64 `json:"user_id"`
}

type guild_message struct {
	Channel_id   uint64               `json:"channel_id"`
	Guild_id     uint64               `json:"guild_id"`
	Message      string               `json:"message"`
	Message_id   string               `json:"message_id"`
	Message_type string               `json:"message_type"`
	Post_type    string               `json:"post_type"`
	Self_id      uint64               `json:"self_id"`
	Self_tiny_id uint64               `json:"self_tiny_id"`
	Sender       guild_message_sender `json:"sender"`
	Sub_type     string               `json:"sub_type"`
	Time         uint64               `json:"time"`
	User_id      uint64               `json:"user_id"`
}

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
		json_text := string(message[:])
		var guild_info guild_message
		json_err := json.Unmarshal([]byte(json_text), &guild_info)
		if json_err != nil {
			fmt.Println(json_err)
		}

		if guild_info.Message_type == "guild" {
			if guild_info.Message == "hello" { //判断收到的消息
				send_message = fmt.Sprintf(`{"action":"send_guild_channel_msg","params":{"guild_id":%d,"channel_id":%d,"message":"%s"}}`, guild_info.Guild_id, guild_info.Channel_id, "Hello World")
				//创建发送消息，可内置函数处理
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
