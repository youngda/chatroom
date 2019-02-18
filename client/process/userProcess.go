package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"chatroom/client/utils"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess)Login(userId int,usePwd string) (err error) {
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

	tf := &utils.Transfer{
		Conn:conn,
	}
    //创建一个tansfor 实例
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("writePkg(conn,data) err = ",err)
		return
	}


	fmt.Printf("消息长度的发送成功=%d,内容=%s",len(data),string(data))
	
	mes,err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn) err = ")
		return
	}
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		//初始化
		CurUser.Conn = conn
		CurUser.User.UserId = userId
		CurUser.User.UserStatus = message.UserOnline

		//显示在线用户列表
		fmt.Println("----在线用户列表如下----")
		for _,v := range loginResMes.UserIds{
			if v == userId{
				continue
			}
			fmt.Println("用户id:\t",v)
			user := &message.User{
				UserId:v,
				UserStatus:message.UserOnline,
			}
			onlineUsers[v] = user
		}
		//fmt.Println("登陆成功-_-")
		go serverProcessMes(conn)
		for true {
			ShowMenu()
		}
	}else if loginResMes.Code == 300 {
		fmt.Println("密码错误")
	}else if loginResMes.Code == 500{
		fmt.Println("用户不存在，请注册")
	}else{
		fmt.Println("服务器未知错误")
	}
	return
}

func (this *UserProcess)Register(userId int,
	userPwd string,userName string)  {

	conn ,err := net.Dial("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("net:Dial err = ",err)
		return
	}
	defer conn.Close()
	//准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//创建一个registerMes 结构体
	var registerMes message.RegisterMes



	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//将regsterMes 序列化
	data , err := json.Marshal(registerMes)
	if err != nil{
		fmt.Println("json:Marshal err=",err)
		return
	}
	//将data 赋值
	mes.Data =string(data)
	//data序列号
	data , err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	tf := &utils.Transfer{
		Conn:conn,
	}
	//创建一个tansfor 实例
	err = tf.WritePkg(data)
	if err != nil{
		fmt.Println("writePkg(conn,data)注册发送消息失败 err = ",err)
		return
	}


	fmt.Printf("消息长度的发送成功=%d,内容=%s",len(data),string(data))

	mes,err = tf.ReadPkg()
	if err != nil{
		fmt.Println("readPkg(conn)注册读取消息 err = ")
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)

	fmt.Println(registerResMes.Code)

	if registerResMes.Code == 201{
		fmt.Println("注册成功-_-")
		os.Exit(0)
	}else if registerResMes.Code == 301 {
		fmt.Println("注册失败...")
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}else{
		fmt.Println("系统错误")
		os.Exit(0)
	}
	return

}
