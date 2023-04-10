package logger

import (
	"fmt"
	"github.com/linkinyy/CycleScan/pkg/types"
	"github.com/linkinyy/CycleScan/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var logger *zap.SugaredLogger

func InitLog() {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志级别使用大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	var encoder zapcore.Encoder
	// 日志格式
	if types.Option.JsonLog {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	var cores []zapcore.Core
	// 是否打印日志
	if !types.Option.NoLog {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel))
	}
	// 是否记录文件
	if types.Option.Record {
		cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(fileWriteSyncer()), zapcore.DebugLevel))
	}
	core := zapcore.NewTee(cores...)
	// zap.AddCaller 添加函数调用信息
	logger = zap.New(core, zap.AddCaller()).Sugar()
}

func fileWriteSyncer() zapcore.WriteSyncer {
	// 在执行目录创建log文件夹, 存放log日志
	logPath := filepath.Join(utils.GetCurrentAbsPath(), "log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		err = os.Mkdir(logPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}
	logFile := filepath.Join(filepath.Join(logPath, "log.log"))
	file, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	return zapcore.AddSync(file)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func DeBug(args ...interface{}) {
	logger.Debug()
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}
