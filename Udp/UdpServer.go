package Udp

import (
	"fmt"
	"log"
	"net"
	"time"
)

const (
	MaxBufferSize int = 65535 //upd 最大发送64kb的包
)

type UdpServer struct {
	udpConn *net.UDPConn
}

func NewUdpServer(ipAddr net.IP, port int) *UdpServer {
	udpServer := new(UdpServer)

	var err error

	// 创建监听
	udpServer.udpConn, err = net.ListenUDP("udp4", &net.UDPAddr{
		IP:   ipAddr,
		Port: port,
	})

	if err != nil {
		udpServer = nil
	}

	return udpServer
}

func (this *UdpServer) clientHandle() {
	// 读取数据
	data := make([]byte, MaxBufferSize)

	//	read, remoteAddr, err := conn.ReadFromUDP(data)

	_, _, err := this.udpConn.ReadFromUDP(data)

	if err != nil {
		log.Fatal(err)
	}
}

func (this *UdpServer) Listen() {
	defer this.udpConn.Close()

	for {
		this.clientHandle()
	}
}

//
//address value format: 127.0.0.1:224.0.1.0
func MulticastListen(address string, h func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println("ResolveUDPAddr error=", err.Error())
		log.Fatal(err)
	}

	l, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("ListenMulticastUDP error=", err.Error())
		log.Fatal(err)
	}

	l.SetReadBuffer(MaxBufferSize)

	for {
		time.Sleep(10 * time.Millisecond)

		buf := make([]byte, MaxBufferSize)
		n, src, err := l.ReadFromUDP(buf)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		h(src, n, buf)
	}
}
