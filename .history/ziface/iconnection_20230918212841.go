package ziface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接的socket
	GetTcpConnection() *net.TCPConn
	// 获取连接 id
	GetConnId() uint32
	// 获取客户端的TCP 状态 IP Port
	Remote() *net.Addr
	//发送数据
	Send(data []byte) error
}

// 定义一个处理连接的方法

type Handler func(*net.TCPConn, []byte, int) error
