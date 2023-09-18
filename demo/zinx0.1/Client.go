package main

import (
	"fmt"
	"net"
	"time"
)

// 模拟客户端

func main1() {

	// 链接远程服务器
	dial, err := net.Dial("tcp4", "127.0.0.1:8999")
	if err != nil {
		return
	}
	for {
		// 写数据
		_, err := dial.Write([]byte("hello zinx0.1"))
		if err != nil {
			return
		}
		buff := make([]byte, 512)
		n, err := dial.Read(buff)
		if err != nil {
			return
		}
		fmt.Printf("server call back : %s conn: %d", buff, n)
		// 阻塞一秒
		time.Sleep(time.Second * 1)
	}

}
