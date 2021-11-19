package functions

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

func GetMessage(Msg string) string {
	json_text := Msg
	send_message := ""
	var guild_info guild_message
	json_err := json.Unmarshal([]byte(json_text), &guild_info)
	if json_err != nil {
		fmt.Println(json_err)
	}
	if guild_info.Message_type == "guild" {
		fmt.Printf("\x1b[%dm %s \x1b[0m\n", 36, "["+strconv.FormatUint(guild_info.Guild_id, 10)+"-"+strconv.FormatUint(guild_info.Channel_id, 10)+"] Sender:"+guild_info.Sender.Nickname+"("+strconv.FormatUint(guild_info.Sender.User_id, 10)+"):"+guild_info.Message)
		if guild_info.Self_tiny_id == guild_info.User_id {
			return ""
		}
		if guild_info.Message == "hello" { //判断收到的消息
			send_message = CreatSendMsg(guild_info.Guild_id, guild_info.Channel_id, "hello")
			//创建发送消息，可内置函数处理
		}
	}
	return send_message
}
func CreatSendMsg(guild_id uint64, channel_id uint64, msg string) string {
	newtext := fmt.Sprintf(`{"action":"send_guild_channel_msg","params":{"guild_id":%d,"channel_id":%d,"message":"%s"}}`, guild_id, channel_id, msg)
	return newtext
}
