package utils

import (
	"os"
)

func GetCurrentAbsPath() string {
	dir, _ := os.Getwd() // 获取当前工作目录
	return dir
}
