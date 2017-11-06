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
	ip = addrs[0].String()
	return
}
