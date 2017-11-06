package main

import (
	"flag"
	"runtime"
	"time"
	"math/rand"
	logger "github.com/shengkehua/xlog4go"
	"net"
	"utils"
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

	utils.BroadCastClient()


}
