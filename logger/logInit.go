package logger

import (
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/skycandyzhe/go-com/config"
	"github.com/skycandyzhe/go-com/file"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// var Logger *zap.SugaredLogger

// var errLvel = zapcore.ErrorLevel
// var infoLevel = zapcore.infoLevel
// var enableConsole = true
// var infoLogPath=""
// var errLogPath=""
var (
	Logger        MyLoggerInterface
	mylogger      *zap.SugaredLogger
	errLvel       = zapcore.ErrorLevel
	infoLevel     = zapcore.InfoLevel
	enableConsole = true
	infoLogPath   = "logs"
	errLogPath    = "logs/err"
	lOG_Name      = "msg"
	mu            sync.Mutex
)

type MyLoggerInterface interface {

	// WithField(key string, value interface{}) *Entry
	// WithFields(fields Fields) *Entry
	// WithError(err error) *Entry

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
	// Printf(format string, args ...interface{})
	// Print(args ...interface{})
	// Debugln(args ...interface{})
	// Infoln(args ...interface{})
	// Println(args ...interface{})
	// Warnln(args ...interface{})
	// Warningln(args ...interface{})
	// Errorln(args ...interface{})
	// Fatalln(args ...interface{})
	// Panicln(args ...interface{})

}

func SetupLogger(log MyLoggerInterface) {
	if log != nil {
		Logger = log
	} else {
		Logger = GetDefaultLogger()
	}
}

func GetDefaultLogger() *zap.SugaredLogger {

	mu.Lock()
	defer mu.Unlock()

	if mylogger != nil {
		return mylogger
	}
	conf := config.GetDefaultConf()
	if conf != nil {
		enableConsole = conf.Console
		if conf.DebugFlag {
			// fmt.Println("need debug log")
			infoLevel = zapcore.DebugLevel
		}
		infoLogPath = conf.Logs.Log_path

		errLogPath = conf.Logs.Err_path
		if infoLogPath == "" {
			infoLogPath = "logs"
		}
		if errLogPath == "" {
			errLogPath = "logs/err"
		}
		lOG_Name = conf.Logs.LogName
		if lOG_Name == "" {
			lOG_Name = "msg"
		}

	}
	if !file.Exists(infoLogPath) {
		file.Mkdir(infoLogPath)
	}
	if !file.Exists(errLogPath) {
		file.Mkdir(errLogPath)
	}
	// ?????????????????????????????? ??????????????????????????????????????????zap??????????????????
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

	// ?????????????????????????????????interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= infoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= errLvel
	})

	// ?????? info???error???????????????io.Writer ?????? getWriter() ???????????????

	infoWriter := getWriter(path.Join(infoLogPath, lOG_Name))
	errorWriter := getWriter(path.Join(errLogPath, lOG_Name))

	// ?????????????????????Logger

	var core zapcore.Core
	if enableConsole {
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
	log := zap.New(core, zap.AddCaller()) // ???????????? zap.AddCaller() ?????????????????????????????????????????????, ????????????
	mylogger = log.Sugar()
	return mylogger
}

func getWriter(filename string) io.Writer {
	// ??????rotatelogs???Logger ???????????????????????? demo.log.YYmmddHH
	// demo.log??????????????????????????????
	// ??????7?????????????????????1??????(??????)??????????????????
	hook, err := rotatelogs.New(
		strings.ReplaceAll(filename, ".log", "") + "-%Y%m%d%H.log",
		//rotatelogs.WithLinkName(filename),
		//rotatelogs.WithMaxAge(time.Hour*24*7),
		//rotatelogs.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(err)
	}
	return hook
}
