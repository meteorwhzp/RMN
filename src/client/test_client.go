package main

import (
	"log"
	"os"
	"github.com/dustin/go-coap"
	"flag"
	"runtime"
	"time"
	"strconv"
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

	ip := net.ParseIP("224.0.0.250")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer conn.Close()

	conn.Write([]byte("hello"))
	logger.Info("<%s>", conn.RemoteAddr())

}
