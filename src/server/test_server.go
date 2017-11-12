package main

import (
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

type LightResource struct {
	m_power string
	m_lightUri string
	m_lightRep resource.Representation
}

func (light *LightResource) put(rep resource.Representation) {
	if ok := rep.GetValue("power", light.m_power); ok == true {
		logger.Info("  power:%v", light.m_power)
	}else {
		logger.Info("power not found in the representation")
	}
}

func (light *LightResource) post(rep resource.Representation) {

}


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
