package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

var onlineUsers map[int]*message.User = make(map[int]*message.User,10)
var CurUser  model.CurUser

//在客户端显示
func outPutOnlineUser()  {
	fmt.Println("---显示在线用户列表---")
	for id,_ := range onlineUsers{
		fmt.Println("用户id:\t",id)
	}
}
func updataUserStatus(notifyUserMes *message.NotifyUserStatusMes){
	user,ok := onlineUsers[notifyUserMes.UserId]
	if !ok{
		user = &message.User{
			UserId:notifyUserMes.UserId,
			UserStatus:notifyUserMes.UserStatus,
		}
	}
	user.UserStatus = notifyUserMes.UserStatus
	onlineUsers[notifyUserMes.UserId] = user
	outPutOnlineUser()
}


