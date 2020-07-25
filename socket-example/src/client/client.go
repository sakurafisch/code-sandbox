package client

import "net"

// SendData returns void
func SendData() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	conn, _ := net.DialTCP("tcp", nil, addr)
	defer conn.Close()
	conn.Write([]byte("客户端发送的数据"))
}
