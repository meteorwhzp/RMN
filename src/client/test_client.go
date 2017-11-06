package main

import (
	"flag"
	"runtime"
	"time"
	"math/rand"
	logger "github.com/shengkehua/xlog4go"
	"net"
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

	/*sip := net.ParseIP("127.0.0.1")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: sip, Port: 5555}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Error(err.Error())
	}
	defer conn.Close()
	conn.Write([]byte("hello"))
	logger.Info("<%s>", conn.RemoteAddr())
*/
	ip := net.ParseIP("224.0.0.250")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer conn.Close()

	data := make([]byte, 1024)
	for {
		conn.Write([]byte("hello"))
		logger.Info("<%s>", conn.RemoteAddr())
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			logger.Error("Error during read: %s", err)
		}
		logger.Info("<%s> %s", remoteAddr, data[:n])
	}

}
