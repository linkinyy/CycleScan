package main

import (
	"github.com/linkinyy/CycleScan/pkg/logger"
	_ "github.com/linkinyy/CycleScan/pkg/types"
)

func main() {
	logger.InitLog()
	logger.Info("This is Info Test")
}
