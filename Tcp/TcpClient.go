package Tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var quitSemaphore chan bool

func (this *TCP) Connect() (*net.TCPConn, error) {
	var (
		tcpAddr *net.TCPAddr
		conn    *net.TCPConn
		err     error
	)

	listenAddress := strings.Join([]string{this.IP, ":", strconv.Itoa(this.Port)}, "")

	tcpAddr, err = net.ResolveTCPAddr("tcp", listenAddress)

	if err != nil {
		return nil, err
	}

	conn, err = net.DialTCP("tcp", nil, tcpAddr)

	return conn, err
}

func (this *TCP) SendStringAndListen(data string) {

	conn, err := this.Connect()

	if err != nil {
		panic(err)
	}

	defer func() {
		fmt.Println("unknow reson close tcp connect")
		this.Close(conn)
	}()

	this.SendString(conn, data)

	go this.onMessageRecived(conn)

	//	<-quitSemaphore
}

func (this *TCP) onMessageRecived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)

	for {
		time.Sleep(10 * time.Millisecond)

		data, err := reader.ReadString('\n')

		if err != nil {
			//			quitSemaphore <- true
			break
		}

		//回调
		this.Callback(conn, data, err)
		break
	}
}
