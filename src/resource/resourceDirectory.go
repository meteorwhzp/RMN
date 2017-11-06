package resource

import (
	"encoding/json"
	logger "github.com/shengkehua/xlog4go"
)

type DiscoveryRes struct{
	IP string `json:"ip"`
	ResourceList []string `json:"resourcelist"`
}

func JsonToDr(b []byte) DiscoveryRes {
	dr := DiscoveryRes{}
	if err := json.Unmarshal(b, &dr); err != nil {
		logger.Error(err.Error())
	}
	return dr
}

func DrToJson(dr DiscoveryRes) []byte {
	jsons, err := json.Marshal(dr)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info(string(jsons))
	return jsons
}

func TestJson() {
	rl := []string{"/a", "/b"}
	dr := DiscoveryRes{
		IP:"123.0.0.1",
		ResourceList: rl,
	}
	DrToJson(dr)
}