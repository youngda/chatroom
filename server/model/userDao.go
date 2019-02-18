package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//服务器启动时，自动创建UserDao 实例
var (
	MyUserDao *UserDao
)
type UserDao struct {
	pool *redis.Pool
}
//使用工程模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool)(userDao *UserDao){
	userDao = &UserDao{
		pool,
	}
	return
}


func (this *UserDao)getUserById (conn redis.Conn,id int) (user *User,err error){
	res,err := redis.String(conn.Do("Hget","users",id))
	if err != nil{
		if err == redis.ErrNil{
			err = ERROR_USER_NOTEXISTS
			return
		}
	}

	user = &User{}
	err = json.Unmarshal([]byte(res),user)

	if err != nil{
		fmt.Println("json.Unmarshal([]byte(res),user)",err)
		return
	}
	return
}
//登陆
//如果id 和密码正确，则返回一个user 实例
func (this *UserDao)Login(userId int,userPwd string)(user *User,err error)  {

	conn := this.pool.Get()
	defer conn.Close()
	user ,err = this.getUserById(conn,userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd{
		err = ERROR_USER_NOTPWD
		return
	}
	fmt.Println("用户信息",userId,user.UserId,user.UserPwd,user.UserName)
	return
}



func (this *UserDao)Register(user *message.User)(err error)  {
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data,err := json.Marshal(user)
	if err != nil{
		return
	}
	_,err = conn.Do("HSet","users",user.UserId,string(data))
	if err != nil{
		fmt.Println("保存用户注册错误 err = ",err)
		return
	}
	return

}
