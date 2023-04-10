package main

import (
	"encoding/json"
	"github.com/linkinyy/CycleScan/pkg/host"
	"github.com/linkinyy/CycleScan/pkg/logger"
	"github.com/linkinyy/CycleScan/pkg/types"
	_ "github.com/linkinyy/CycleScan/pkg/types"
)

func main() {
	logger.InitLog()
	target := host.Target{
		Ip:    types.Option.Ip,
		Ports: make([]host.Port, 0, 10),
		Os:    make([]string, 0, 3),
	}
	target.Scan(types.Option.Ports.Value())
	if !target.IsAlive() {
		logger.Error("Target Is Not Alive")
		logger.Error("Exit Scan!!!")
		return
	}
	marshal, _ := json.Marshal(target)
	logger.Info(string(marshal))
}
