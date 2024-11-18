package goos

import (
	"errors"
	"net"
	"runtime"
	"strings"
)

// determine whether it is a Windows system
func IsWindows() bool {
	return strings.EqualFold(runtime.GOOS, "windows")
}

func IsMac() bool {
	return strings.EqualFold(runtime.GOOS, "darwin")
}

// IsLinux 判断当前操作系统是否为Linux系统
func IsLinux() bool {
	return strings.EqualFold(runtime.GOOS, "linux")
}

// Obtain MAC address
func GetMACAddress() (string, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	var mac string
	var bakMac1 string
	var bakMac2 string

	for i := 0; i < len(netInterfaces); i++ {
		flags := netInterfaces[i].Flags.String()
		if strings.Contains(flags, "up") &&
			strings.Contains(flags, "broadcast") &&
			strings.Contains(flags, "running") &&
			!strings.Contains(flags, "loopback") {

			if strings.Contains(netInterfaces[i].Name, "WLAN") {
				mac = netInterfaces[i].HardwareAddr.String()
				return mac, nil
			}
			// 感谢@lazyphp 优化windows系统下，启用了WSL系统导致MAC获取地址变化的问题
			if !strings.Contains(netInterfaces[i].Name, "VMware") && !strings.Contains(netInterfaces[i].Name, "WSL") {
				bakMac1 = netInterfaces[i].HardwareAddr.String()
			} else {
				bakMac2 = netInterfaces[i].HardwareAddr.String()
			}
		}
	}
	if mac == "" {
		if bakMac1 != "" {
			return bakMac1, nil
		}
		return bakMac2, nil
	}
	return mac, errors.New("unable to get the correct MAC address")
}
