## 在线聊天系统 Golang开发

### 已完成完成

* 用户注册
* 用户登陆
* 消息群发
* 在线用户查询

### 待完成

* 用户退出状态维护
* 点对点聊天
* 界面优化

## 环境依赖

* 包： "github.com/garyburd/redigo/redis"
* 运行环境： redis 服务器，默认端口无需改动

## windows端测试(linux方法相仿)

* 安装redis golang插件，cmd 中运行指令：go get github.com/garyburd/redigo/redis

* 下载程序到golang项目src目录(如G:\go\src)

* 进入G:\go\src\chatroom\Redis-x64-3.2.100，点击redis-server.exe 运行redis服务端

* 在cmd中执行指令生成可执行程序(src目录)：
>> `go build -o chatroom/server.exe chatroom/server/main`
>> `go build -o chatroom/client.exe chatroom/client/main`
>> `start chatroom/server.exe`
>> `start chatroom/client.exe`

* 如需测试群发则启动多个客户端
>> `start chatroom/client.exe`
>> `start chatroom/client.exe`
>> `start chatroom/client.exe`