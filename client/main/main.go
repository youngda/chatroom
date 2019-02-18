package main

import (
	"fmt"
	"os"
	"chatroom/client/process"
)

var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int


	for true {
		fmt.Println("------------------欢迎登陆聊天室-----------------")
		fmt.Println("\t\t\t1.登陆聊天室")
		fmt.Println("\t\t\t2.用户注册")
		fmt.Println("\t\t\t3.退出")
		fmt.Println("\t\t\t4.请选择（1-3）：")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)

			up := &process.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("----用户注册----")
			fmt.Println("请输入用户ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名:")
			fmt.Scanf("%s\n", &userName)

			up := &process.UserProcess{}
			up.Register(userId,userPwd,userName)
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("\t\t\t1.请选择（1-3）：")

		}
	}


}
