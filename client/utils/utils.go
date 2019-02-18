package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg()(mes message.Message,err error)  {

	fmt.Println("读取客户端发送的数据...")
	_,err = this.Conn.Read(this.Buf[0:4])
	if err != nil{
		//fmt.Println("conn.Read err= ",err)
		return
	}
	//fmt.Println("读到的buf = ",buf[:4])

	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen 读取消息内容

	n,err := this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil{
		fmt.Println("conn.Read(buf[:pkglen]) err = ")
		return
	}
	//把pkgLen 反序列化->message.Message
	err = json.Unmarshal(this.Buf[:pkglen],&mes)
	if err != nil{
		fmt.Println("json.Unmarshal(buf[:pkglen],&mes) err = ")
		return
	}
	return
}
func (this *Transfer) WritePkg(data []byte)(err error)  {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)
	//发送长度
	n,err := this.Conn.Write(buf[:4])
	if n !=4 || err != nil{
		fmt.Println("conn.Write(len) err=",err)
		return
	}
	//发送消息体
	n,err = this.Conn.Write(data)
	if n !=int(pkgLen) || err != nil{
		fmt.Println("conn.Write(data) err=",err)
		fmt.Println("server err")
		return
	}
	return
}
