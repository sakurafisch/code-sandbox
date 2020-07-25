package server

import (
	"fmt"
	"net"
)

// RunServer returns void
func RunServer() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899") // 创建服务器地址
	lis, _ := net.ListenTCP("tcp4", addr)                   // 创建监听器
	fmt.Println("服务器已经启动")
	conn, _ := lis.Accept() // 通过监听器获取客户端传递过来的数据
	b := make([]byte, 1024) // 转换数据
	n, _ := conn.Read(b)
	defer conn.Close()
	fmt.Println("获取到的数据为：", string(b[:n]))
}
