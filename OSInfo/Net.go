package OSInfo

import (
	"net"

	thrNet "github.com/shirou/gopsutil/net"
)

type Net struct{}

var (
	nt thrNet.IOCountersStat
)

func init() {
	nts, _ := thrNet.IOCounters(true)

	if nts != nil && len(nts) > 0 {
		nt = nts[0]
	}
}

//网络接收字节数，单位：字节(byte)
func (n *Net) BytesRecv() uint64 {
	return nt.BytesRecv
}

//网络发送字节数，单位：字节(byte)
func (n *Net) BytesSend() uint64 {
	return nt.BytesSent
}

//服务器node所有ipv4地址
func (n *Net) IPAddress() []net.IP {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return nil
	}

	var netIPArray []net.IP

	netIPArray = make([]net.IP, 0, 0)

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				netIPArray = append(netIPArray, ipnet.IP)
			}
		}
	}

	if len(netIPArray) <= 0 {
		netIPArray = nil
	}

	return netIPArray
}
