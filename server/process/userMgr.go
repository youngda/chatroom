package process2

import "fmt"

var (
	userMgr *UserMgr
)
type UserMgr struct {
	onlineUsers map[int] *UserProcess
}

func init()  {
	userMgr = &UserMgr{
		onlineUsers:make(map[int] *UserProcess),
	}
}

func (this *UserMgr)addOnlineUser(up *UserProcess)  {
	this.onlineUsers[up.UserId] = up
}
func (this *UserMgr)DelOnlineUser(userId int)  {
	delete(this.onlineUsers,userId)
}
func (this *UserMgr)GetOnlineById(userId int)(up *UserProcess,err error)  {
	up,ok := this.onlineUsers[userId]
	if !ok{
		err = fmt.Errorf("用户 %d 不存在",userId)
		return
	}
	return
}