package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"chatroom/client/utils"
)

func ShowMenu()  {
	fmt.Println("--------恭喜登陆成功--------")
	fmt.Println("--------1. 显示在线用户列表--------")
	fmt.Println("--------2. 发送消息--------")
	fmt.Println("--------3. 信息列表--------")
	fmt.Println("--------4. 退出系统--------")
	var key int
	var content string

	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
	case 1:
		fmt.Println("显示在线用户列表--")
		outPutOnlineUser()
	case 2:
		fmt.Println("发送消息")
		fmt.Println("请输入你想对大家说的话")
		fmt.Scanf("%s\n",&content)
		smsProcess.SendGroupMes(content)
	case 3:
		fmt.Println("消息列表")
	case 4:
		fmt.Println("你选择退出了系统....")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不对")
	}

}

//和服务器保持通讯
func serverProcessMes(conn net.Conn)  {
	tf :=&utils.Transfer{
		Conn:conn,
	}
	for{
		fmt.Println("客户端正在等待服务器发送的消息")
		mes ,err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg() err = ")
			return
		}
		//如果读取到消息
		//fmt.Printf("mes = %v",mes)
		switch mes.Type {
		case message.NotifyUserStatusMesType:
			//1.取出消息
			//2.把用户信息状态保存到客户端(map)
			//处理
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			updataUserStatus(&notifyUserStatusMes)
		case message.SmsMesType:
			outPutGroupMes(&mes)
		default:
			fmt.Println("服务器返回了未知的数据类型")
			
		}
	}
}