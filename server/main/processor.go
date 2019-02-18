package main

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
	"chatroom/server/process"
)

type Processor struct {
	Conn net.Conn
}


func  (this *Processor)serverProcessMes(mes *message.Message)(err error)  {

	fmt.Println("mes = ",mes)
	switch mes.Type {
	case message.LoginMesType:

		up := &process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		up := &process2.UserProcess{
			Conn:this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		fmt.Println("SmsMes = ",mes)
		smsprocess := &process2.SmsProcess{}
		smsprocess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理")
	}
	return
}
func (this *Processor)Process2() (err error){

	for true {
		tf := &utils.Transfer{
			Conn:this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil{
			//fmt.Println("readPkg(conn) err= ",err)
			if err == io.EOF{
				fmt.Println("客户端退出，服务器也推出")
				return err
			}else{
				fmt.Println("客户端退出，服务器也推出")
				return err
			}

		}
		//fmt.Println("mes = ",mes)
		err = this.serverProcessMes(&mes)
		if err != nil{
			return err
		}
	}
	return

}
