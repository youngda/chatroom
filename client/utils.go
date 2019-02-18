package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"chatroom/common/message"
	"net"
)

func readPkg(conn net.Conn)(mes message.Message,err error)  {
	buf := make([]byte,8096)
	fmt.Println("读取客户端发送的数据...")
	_,err = conn.Read(buf[0:4])
	if err != nil{
		//fmt.Println("conn.Read err= ",err)
		return
	}
	//fmt.Println("读到的buf = ",buf[:4])

	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen 读取消息内容

	n,err := conn.Read(buf[:pkglen])
	if n != int(pkglen) || err != nil{
		fmt.Println("conn.Read(buf[:pkglen]) err = ")
		return
	}
	//把pkgLen 反序列化->message.Message
	err = json.Unmarshal(buf[:pkglen],&mes)
	if err != nil{
		fmt.Println("json.Unmarshal(buf[:pkglen],&mes) err = ")
		return
	}
	return
}
func writePkg( conn net.Conn ,data []byte)(err error)  {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(buf[:4])
	if n !=4 || err != nil{
		fmt.Println("conn.Write(len) err=",err)
		return
	}
	//发送消息体
	n,err = conn.Write(data)
	if n !=int(pkgLen) || err != nil{
		fmt.Println("conn.Write(data) err=",err)
		fmt.Println("what err")
		return
	}
	return
}