package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)



func process(conn net.Conn)  {
	//读客户端发送的信息
	defer conn.Close()
	processor :=&Processor{
		Conn:conn,
	}
	err := processor.Process2()
	if err != nil{
		fmt.Println("客户端和服务器通讯协议错误err=",err)
		return
	}


}
//完成UserDao初始化
func initUserDao()  {
	model.MyUserDao = model.NewUserDao(pool)
}
func main()  {
	initPool("127.0.0.1:6379",8,0,300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口监听....")
	listen , err := net.Listen("tcp","0.0.0.0:8889")
	if err != nil{
		fmt.Println("net.Listen err = ",err)
		return
	}
	//一旦监听成功
	for true {
		fmt.Println("等待客户端连接....")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("listen Accept err = ",err)
		}

		go process(conn)
	}
}
