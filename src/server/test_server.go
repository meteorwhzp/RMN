package main

import (
	"net"
	"flag"
	logger "github.com/shengkehua/xlog4go"
	"math/rand"
	"runtime"
	"time"
	"utils"
	"resource"
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

	resource.TestJson()

	logger.Info("the local ip is %s", utils.GetLocalIP())

	utils.BroadCastServer()

}
