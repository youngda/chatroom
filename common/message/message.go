package message

const (
	LoginMesType            = "LoginMes"
	LoginResMesType         = "LoginResMes"
	RegisterResMesType      = "RegisterResMes"
	RegisterMesType         = "RegisterMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
	)
//定义状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string
	Data string
}

type LoginMes struct {
	UserId int
	UserPwd string
	UserName string
}

type LoginResMes struct {
	Code int
	UserIds []int
	Error string
}
type RegisterMes struct {
	User User
}
type RegisterResMes struct {
	Code int
	Error string
}

//用户状态变化的类型

type NotifyUserStatusMes struct {
	UserId int //用户ID
	UserStatus int //用户状态
}




//发送消息
type SmsMes struct {
	Content string
	User User
}