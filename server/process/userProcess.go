package process2

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}
//通知用户上线
func (this *UserProcess)NotifyOtherOnlineStatus (userId int) {
	//遍历onlineUsers
	for id,up := range userMgr.onlineUsers {
		if id == userId{
			continue
		}
		//开始通知
		up.NotifyOtherOnline(userId)
	}
	return
}
func (this *UserProcess)NotifyOtherOnline(userId int)  {
	//组装消息
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.UserStatus = message.UserOnline
	//序列化第一场层
	data,err := json.Marshal(notifyUserStatusMes)
	if err != nil{
		fmt.Println("notifyUserStatusMes err = ",err)
		return
	}
	//序列化第二层
	mes.Data = string(data)
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal(mes) err = ",err)
		return
	}

	//发送，创建transf实例

	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("tf.WritePkg(data) err = ",err)
		return
	}
	return
}

func (this *UserProcess)ServerProcessRegister(mes *message.Message)(err error)  {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil{
		fmt.Println("json.Unmarshal err=" ,err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesType
	//再声明一个LoginResMes,并完成赋值
	var registerResMes message.RegisterResMes

	//使用
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil{
		if err == model.ERROR_USER_EXISTS{
			registerResMes.Code = 301
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		}else {
			registerResMes.Code = 501
			registerResMes.Error = "注册未知错误..."
		}
	}else {
		registerResMes.Code = 201
	}

	data,err := json.Marshal(registerResMes)
	if err != nil{
		fmt.Println("json.Marshal(loginResMes) err = ",err)
		return
	}
	resMes.Data = string(data)

	data,err = json.Marshal(resMes)

	if err != nil {
		fmt.Println("json.Marshal(resMes) err= ",err)
		return
	}
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)


	return
}

func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error)  {
	var loginmes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginmes)
	if err != nil{
		fmt.Println("json.Unmarshal err=" ,err)
		return
	}
	//先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	//使用
	user,err := model.MyUserDao.Login(loginmes.UserId,loginmes.UserPwd)
	if err != nil{
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == model.ERROR_USER_NOTPWD{
			loginResMes.Code = 400
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code = 505
			loginResMes.Error = "服务器未知错误"
		}

	}else {
		loginResMes.Code = 200
		this.UserId = loginmes.UserId
		userMgr.addOnlineUser(this)

		this.NotifyOtherOnlineStatus(loginmes.UserId)
		//将当前在线用户，放入loginResMes
		for id,_ := range userMgr.onlineUsers{
			loginResMes.UserIds = append(loginResMes.UserIds,id)
		}
		fmt.Println("登陆成功",user)
	}


	data,err := json.Marshal(loginResMes)
	if err != nil{
		fmt.Println("json.Marshal(loginResMes) err = ",err)
		return
	}
	resMes.Data = string(data)

	data,err = json.Marshal(resMes)

	if err != nil {
		fmt.Println("json.Marshal(resMes) err= ",err)
		return
	}
	tf := &utils.Transfer{
		Conn:this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
