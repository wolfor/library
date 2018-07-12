package Tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func (this *TCP) Listen() {
	listenAddress := strings.Join([]string{this.IP, ":", strconv.Itoa(this.Port)}, "")

	var tcpAddr *net.TCPAddr

	tcpAddr, _ = net.ResolveTCPAddr("tcp", listenAddress)

	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)

	defer tcpListener.Close()

	for {
		time.Sleep(10 * time.Millisecond)

		tcpConn, err := tcpListener.AcceptTCP()

		if err != nil {
			continue
		}

		fmt.Println("A client connected : " + tcpConn.RemoteAddr().String())
		go this.tcpPipe(tcpConn)
	}
}

func (this *TCP) tcpPipe(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)

	for {
		time.Sleep(10 * time.Millisecond)

		data, err := reader.ReadString('\n')

		if err != nil {
			continue
		}

		//回调
		this.Callback(conn, data, err)
	}
}
