package logger

import (
	"io"
	"os"
	"path"
	"strings"
	"time"

	"github.com/skycandyzhe/go-com/config"
	"github.com/skycandyzhe/go-com/file"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var Logger *zap.SugaredLogger

// var ErrLvel = zapcore.ErrorLevel
// var InfoLevel = zapcore.InfoLevel
// var EnableConsole = true
// var InfoLogPath=""
// var ErrLogPath=""
var (
	Logger        *zap.SugaredLogger
	ErrLvel       = zapcore.ErrorLevel
	InfoLevel     = zapcore.InfoLevel
	EnableConsole = true
	InfoLogPath   = "logs"
	ErrLogPath    = "logs/err"
	LOG_Name      = "msg"
)

func init() {

	if config.Conf != nil {
		EnableConsole = config.Conf.Console
		if config.Conf.DebugFlag {
			// fmt.Println("need debug log")
			InfoLevel = zapcore.DebugLevel
		}
		InfoLogPath = config.Conf.Logs.Log_path

		ErrLogPath = config.Conf.Logs.Err_path
		if InfoLogPath == "" {
			InfoLogPath = "logs"
		}
		if ErrLogPath == "" {
			ErrLogPath = "logs/err"
		}
		LOG_Name = config.Conf.Logs.LogName
		if LOG_Name == "" {
			LOG_Name = "msg"
		}

	}
	if !file.Exists(InfoLogPath) {
		file.Mkdir(InfoLogPath)
	}
	if !file.Exists(ErrLogPath) {
		file.Mkdir(ErrLogPath)
	}
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
		return lvl >= InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= ErrLvel
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现

	infoWriter := getWriter(path.Join(InfoLogPath, LOG_Name))
	errorWriter := getWriter(path.Join(ErrLogPath, LOG_Name))

	// 最后创建具体的Logger

	var core zapcore.Core
	if EnableConsole {
		// fmt.Print("need console ")
		core = zapcore.NewTee(
			//zapcore.NewCore(zapcore.NewConsoleEncoder(enConfig), zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel),
		)

	} else {
		core = zapcore.NewTee(
			//zapcore.NewCore(zapcore.NewConsoleEncoder(enConfig), zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
			zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		)
	}
	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	Logger = log.Sugar()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		strings.ReplaceAll(filename, ".log", "") + "-%Y%m%d%H.log", // 没有使用go风格反人类的format格式
		//rotatelogs.WithLinkName(filename),
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
