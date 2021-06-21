package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"powershellDeal/common/config"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func init() {
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
		return lvl >= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现

	infoWriter := getWriter(config.Conf.Logs.Log_path)
	errorWriter := getWriter(config.Conf.Logs.Err_path)

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		//zapcore.NewCore(zapcore.NewConsoleEncoder(enConfig), zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), infoLevel),
	)

	log := zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	Logger = log.Sugar()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	hook, err := rotatelogs.New(
		strings.Replace(filename, ".log", "", -1) + "-%Y%m%d%H.log", // 没有使用go风格反人类的format格式
		//rotatelogs.WithLinkName(filename),
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// func Debug(args ...interface{}) {
// 	Logger.Debug(args...)
// }

// func Debugf(template string, args ...interface{}) {
// 	Logger.Debugf(template, args...)
// }

// func Info(args ...interface{}) {
// 	Logger.Info(args...)
// }

// func Infof(template string, args ...interface{}) {
// 	Logger.Infof(template, args...)
// }

// func Warn(args ...interface{}) {
// 	Logger.Warn(args...)
// }

// func Warnf(template string, args ...interface{}) {
// 	Logger.Warnf(template, args...)
// }

// func Error(args ...interface{}) {
// 	Logger.Error(args...)
// }

// func Errorf(template string, args ...interface{}) {
// 	Logger.Errorf(template, args...)
// }

// func DPanic(args ...interface{}) {
// 	Logger.DPanic(args...)
// }

// func DPanicf(template string, args ...interface{}) {
// 	Logger.DPanicf(template, args...)
// }

// func Panic(args ...interface{}) {
// 	Logger.Panic(args...)
// }

// func Panicf(template string, args ...interface{}) {
// 	Logger.Panicf(template, args...)
// }

// func Fatal(args ...interface{}) {
// 	Logger.Fatal(args...)
// }

// func Fatalf(template string, args ...interface{}) {
// 	Logger.Fatalf(template, args...)
// }
