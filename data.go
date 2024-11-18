package machine_code

import (
	"github.com/super-l/machine-code/types"
	"net"
	"strings"
)

var Machine types.MachineInformation
var MachineErr error

// Obtain accurate export traffic IP information
func GetIpAddr() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return "", err
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip := strings.Split(localAddr.String(), ":")[0]
	return ip, nil
}

// Get all IP addresses
func GetIpAddrAll() ([]string, error) {
	var ipList []string
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		return ipList, err
	}
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && !ipNet.IP.IsLinkLocalUnicast() {
			if ipNet.IP.To4() != nil {
				ipList = append(ipList, ipNet.IP.To4().String())
			}
		}
	}
	return ipList, nil
}
