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

	// 获取客户端的TCP 状态 IP Port

	//发送数据
}
