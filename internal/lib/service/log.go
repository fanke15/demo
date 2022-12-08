package service

import (
	"cron/internal/lib/define"
	"cron/pkg"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultSuffix = ".log"
)

var (
	// zap操作对象
	zapLog            *zap.Logger
	defaultLogDirName = "logfiles"
	defaultKV         = make([]interface{}, define.Zero)
)

type Log struct {
	zapLog *zap.Logger
}

// GetLog 项目初始化
func GetLog() *Log {
	return &Log{zapLog}
}

func (l *Log) Info(errMsg string, others ...interface{}) {
	filed := customKV(others...)
	l.zapLog.Info(errMsg, filed...)
}

func (l *Log) Error(errMsg string, others ...interface{}) {
	filed := customKV(others...)
	l.zapLog.Error(errMsg, filed...)
}

//---------------------------内部私有方法---------------------------//

func initLog() {
	var (
		// 文件存储设置 输出hook
		hook = lumberjack.Logger{
			Filename:   pkg.AnySliceToStr(define.Forwardslash, defaultServiceDir, defaultLogDirName, pkg.AnySliceToStr(define.StrNull, time.Now().Format(define.DateFormat4), defaultSuffix)),
			MaxSize:    define.Ten,
			MaxBackups: define.Five, // 文件数过多时进行清理
			MaxAge:     define.One,  // 切割间隔天数
			Compress:   false,
		}
		// 编码器配置,zap日志初始化设置
		logEncodeConfig = zapcore.EncoderConfig{
			TimeKey:       "time",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			// EncodeLevel:   zapcore.LowercaseLevelEncoder, // 小写编码器
			EncodeLevel: zapcore.CapitalColorLevelEncoder, // 色彩
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format(define.DateFormat1))
			},
			EncodeDuration:   zapcore.SecondsDurationEncoder,
			EncodeCaller:     zapcore.ShortCallerEncoder, // 半路径编码器
			ConsoleSeparator: define.StrSpace,
		}

		// 日志输出级别设置
		tree = []zapcore.Core{
			zapcore.NewCore(
				zapcore.NewConsoleEncoder(logEncodeConfig),
				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), // 输出到控制台os.Stdout
				zapcore.InfoLevel,
			)}
	)
	logEncodeConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	tree = append(tree, zapcore.NewCore( // 输出到文件hook
		zapcore.NewJSONEncoder(logEncodeConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
		zapcore.ErrorLevel,
	))

	core := zapcore.NewTee(tree...)
	zapLog = zap.New(core,
		zap.AddCallerSkip(define.One), zap.WithCaller(true),
		zap.Fields(customKV(defaultKV...)...)) // 构造日志
}

// 自定义默认k-v值
func customKV(kvs ...interface{}) []zap.Field {
	var (
		num = len(kvs) / define.Two
		fs  []zap.Field
	)
	if num > 0 {
		for i := define.One; i <= num; i++ {
			f := zap.Any(pkg.AnyToStr(kvs[define.Two*i-define.Two]), kvs[define.Two*i-define.One])
			fs = append(fs, f)
		}
	}
	return fs
}
