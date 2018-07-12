// Tcp project Tcp.go
package Tcp

import (
	"bufio"
	"fmt"
	"net"
)

type RecvCallback func(*net.TCPConn, string, error)

type TCP struct {
	IP       string
	Port     int
	Callback RecvCallback
}

func NewTcp(ip string, port int, callback RecvCallback) *TCP {
	tcp := new(TCP)
	tcp.IP = ip
	tcp.Port = port
	tcp.Callback = callback

	return tcp
}

func (this *TCP) Close(conn *net.TCPConn) {
	ipStr := conn.RemoteAddr().String()
	fmt.Println("disconnected :" + ipStr)

	conn.Close()
}

func (this *TCP) SendByteArray(conn *net.TCPConn, data []byte) {
	conn.Write(data)
}

func (this *TCP) SendString(conn *net.TCPConn, data string) {
	buf := []byte(data)

	this.SendByteArray(conn, buf)
}

func (this *TCP) SendStringAndRecvCallback(conn *net.TCPConn, data string, callback RecvCallback) {
	buf := []byte(data)

	this.SendByteArray(conn, buf)

	go this.onStringRecv(conn)
}

func (this *TCP) onStringRecv(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		//回调
		this.Callback(conn, data, err)
		break
	}
}
