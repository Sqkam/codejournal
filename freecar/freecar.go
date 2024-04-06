package freecar

import (
	"fmt"
	"net"
)

func main() {
	GetIP()
	_, _ = GetOutboundIP()
}
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "[2001:db8::1]:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}

func GetIP() {

	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, _ := i.Addrs()

		fmt.Println(i.Name)
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Printf("%+v\n", ip)
		}
	}

}
