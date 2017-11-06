package utils

import (
	"net"
	logger "github.com/shengkehua/xlog4go"
)

func GetLocalIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				logger.Info(ipnet.IP.String())
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return
}
