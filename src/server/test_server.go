package main

import (
	"net"
	"flag"
	logger "github.com/shengkehua/xlog4go"
	"math/rand"
	"runtime"
	"time"
)

var (
	logconf = flag.String("l", "./conf/log.json", "log config file path")
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())
	flag.Parse()

	// log
	if err := logger.SetupLogWithConf(*logconf); err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("start server")

	listener, err:= net.ListenUDP("udp", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5555})
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Local: <%s>", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			logger.Error(err.Error())
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			logger.Error(err.Error())
		}
	}

	/*addr, err := net.ResolveUDPAddr("udp", "224.0.0.250:9981")
	if err != nil {
		logger.Error(err.Error())
		return
	}

	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	logger.Info("Local: <%s>", listener.LocalAddr().String())
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			logger.Error("Error during read: %s", err)
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
	}*/
}
