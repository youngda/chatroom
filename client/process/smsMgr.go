package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outPutGroupMes(mes *message.Message)  {

	//反序列化....
	var smsMes message.SmsMes
	err  := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal([]byte(mes.Data),&smsMes) err = ",err)
		return
	}
	info := fmt.Sprintf("用户ID:\t%d 对大家说：\t %s",smsMes.User.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
