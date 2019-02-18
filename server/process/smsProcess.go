package process2

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {
	
}

func (this *SmsProcess)SendGroupMes(mes *message.Message)  {
	//遍历服务器map,将消息转发
	var smsMes message.SmsMes
	err :=json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("json.Unmarshal([]byte(mes.Data),&smsMes) err = ",err)
		return
	}
	data,err := json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal(mes) err = ",err)
		return
	}
	for id,up := range userMgr.onlineUsers{
		if id == smsMes.User.UserId{
			continue
		}
		this.SendMesToEachOnlineUser(data,up.Conn)
	}
}
func (this *SmsProcess)SendMesToEachOnlineUser(data []byte,conn net.Conn)  {
	tf := &utils.Transfer{
		Conn:conn,
	}
	err := tf.WritePkg(data)
	if err != nil{
		fmt.Println("转发消息出错tf.WritePkg(data) err =",err)
	}
}