package Udp

import (
	"net"
)

type UdpClient struct {
	udpConn *net.UDPConn
}

func NewUdpClient(ipAddr net.IP, port int) *UdpClient {
	udpClient := new(UdpClient)

	var err error

	// 创建连接
	udpClient.udpConn, err = net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   ipAddr,
		Port: port,
	})

	if err != nil {
		return nil
	}

	return udpClient
}

func (this *UdpClient) Send(buf []byte) bool {
	var isSucess bool

	if this.udpConn == nil {
		return false
	}

	if buf == nil || len(buf) <= 0 {
		return false
	}

	defer this.udpConn.Close()

	_, err := this.udpConn.Write(buf)
	if err != nil {
		return false
	}

	isSucess = true

	return isSucess
}
