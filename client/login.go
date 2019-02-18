package main

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int,usePwd string) (err error) {
	//fmt.Printf("userID = %d userPwd = %s\n",userId,usePwd)


	//return nil
	conn ,err := net.Dial("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("net:Dial err = ",err)
		return
	}
	defer conn.Close()


	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = usePwd


	//将loginMess 序列化
	data , err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json:Marshal err=",err)
		return
	}
	//将data 赋值
	mes.Data =string(data)
	data , err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}


	err = writePkg(conn,data)
	if err != nil{
		fmt.Println("writePkg(conn,data) err = ",err)
		return
	}


	fmt.Printf("消息长度的发送成功=%d,内容=%s",len(data),string(data))


    mes,err = readPkg(conn)
    if err != nil{
    	fmt.Println("readPkg(conn) err = ")
    	return
	}
    var loginResMes message.LoginResMes
    err = json.Unmarshal([]byte(mes.Data),&loginResMes)
    if loginResMes.Code == 200{
    	fmt.Println("登陆成功-_-")
	}else if loginResMes.Code == 300 {
		fmt.Println("密码错误")
	}else {
		fmt.Println("用户不存在，请注册")
	}
	return

}
