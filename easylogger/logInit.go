package easylogger

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/skycandyzhe/go-com/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger        LoggerINF
	mylogger      *zap.SugaredLogger
	infoLevel     = zapcore.InfoLevel
	enableConsole = true
	infoLogPath   = "logs"
	log_name      = "msg"
	mu            sync.Mutex
)

type LoggerINF interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

func Init() {
	Logger = GetDefaultLogger()
}

func SetupOtherLogger(log LoggerINF) {
	if log != nil {
		Logger = log
	} else {
		Logger = GetDefaultLogger()
	}
}
func GetLogger(cnfpath string) *zap.SugaredLogger {

	mu.Lock()
	defer mu.Unlock()

	if mylogger != nil {
		return mylogger
	}
	conf := config.GetConf(cnfpath)
	enableConsole = conf.Console
	if conf.DebugFlag {
		// fmt.Println("need debug log")
		infoLevel = zapcore.DebugLevel
	}
	infoLogPath = conf.LogPath

	log_name, _ := os.Executable()
	_, log_name = filepath.Split(log_name)
	log_name = strings.TrimSuffix(log_name, path.Ext(log_name))

	os.Mkdir(infoLogPath, os.ModePerm)
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= infoLevel
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(path.Join(infoLogPath, log_name))

	// 最后创建具体的Logger
	var core zapcore.Core
	if enableConsole {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel),
		)
	} else {
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		)
	}
	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	mylogger = log.Sugar()
	return mylogger
}
func GetDefaultLogger() *zap.SugaredLogger {
	return GetLogger("log_config.yaml")
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每24H分割一次日志,单个日志最大20M 最多保留7个日志
	if runtime.GOOS == "windows" {
		hook, err := rotatelogs.New(
			strings.ReplaceAll(filename, ".log", "")+"-%Y%m%d%H%M.log",
			rotatelogs.WithRotationTime(time.Hour*24), //rotate 最小为5分钟轮询。默认60s  低于1分钟就按1分钟来
			rotatelogs.WithRotationSize(20*1024*1024),
			rotatelogs.WithRotationCount(7),
		)
		if err != nil {
			panic(err)
		}
		return hook
	}
	hook, err := rotatelogs.New(
		strings.ReplaceAll(filename, ".log", "")+"-%Y%m%d%H%M.log",
		rotatelogs.WithLinkName(filename+"_latest.log"), // 生成软链，指向最新日志文件
		rotatelogs.WithRotationTime(time.Hour*24),       //rotate 最小为5分钟轮询。默认60s  低于1分钟就按1分钟来
		rotatelogs.WithRotationSize(20*1024*1024),
		rotatelogs.WithRotationCount(7),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
