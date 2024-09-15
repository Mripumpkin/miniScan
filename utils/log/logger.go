package log

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	config "miniScan/utils/conf"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// 定义日志实例
var defaultLogger *logrus.Logger
var LogPortScan *logrus.Logger
var LogDomain *logrus.Logger
var LogWebInfo *logrus.Logger
var logger *logrus.Logger
var once sync.Once

func init() {
	// 初始化不同的日志实例
	LogPortScan = newLogrusLoggerWithRotation(config.Config(), "portscan", "端口扫描")
	LogDomain = newLogrusLoggerWithRotation(config.Config(), "domain", "子域名扫描")
	LogWebInfo = newLogrusLoggerWithRotation(config.Config(), "web", "web探测")
	defaultLogger = newLogrusLogger(config.Config())
}

// NewLogger returns a configured logrus instance
func NewLogger(cfg config.Provider) *logrus.Logger {
	once.Do(func() {
		logger = newLogrusLogger(cfg)
	})
	return logger
}

// newLogrusLogger 创建基本的日志实例
func newLogrusLogger(cfg config.Provider) *logrus.Logger {
	l := logrus.New()
	if cfg.GetBool("json_logs") {
		l.Formatter = new(logrus.JSONFormatter)
	}
	l.Out = os.Stderr

	switch cfg.GetString("loglevel") {
	case "debug":
		l.Level = logrus.DebugLevel
	case "warning":
		l.Level = logrus.WarnLevel
	case "info":
		l.Level = logrus.InfoLevel
	default:
		l.Level = logrus.DebugLevel
	}
	return l
}

// newLogrusLoggerWithRotation 创建带有日志轮转功能的日志实例
func newLogrusLoggerWithRotation(cfg config.Provider, filename string, prefix string) *logrus.Logger {
	l := logrus.New()

	// 设置文件名，包含功能名称和日期
	date := time.Now().Format("2006-01-02")
	logFile := filepath.Join("logs", filename+"-"+date+".log")

	// 使用 lumberjack 控制日志大小和轮转
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1,    // 文件最大10MB
		MaxBackups: 10,   // 保留最近3个备份
		MaxAge:     30,   // 日志保存最大28天
		Compress:   true, // 是否压缩旧日志文件
	}

	l.Out = lumberjackLogger

	// 根据配置设置日志级别
	switch cfg.GetString(prefix + ".Level") {
	case "debug":
		l.Level = logrus.DebugLevel
	case "warning":
		l.Level = logrus.WarnLevel
	case "info":
		l.Level = logrus.InfoLevel
	default:
		l.Level = logrus.DebugLevel
	}

	// 设置日志前缀
	l.WithField("prefix", prefix)

	return l
}

// Fields is a map string interface to define fields in the structured log
type Fields map[string]interface{}

// With allow us to define fields in out structured logs
func (f Fields) With(k string, v interface{}) Fields {
	f[k] = v
	return f
}

// WithFields allow us to define fields in out structured logs
func (f Fields) WithFields(f2 Fields) Fields {
	for k, v := range f2 {
		f[k] = v
	}
	return f
}

// Debug package-level convenience method.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf package-level convenience method.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Debugln package-level convenience method.
func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

// Error package-level convenience method.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf package-level convenience method.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Errorln package-level convenience method.
func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

// Fatal package-level convenience method.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf package-level convenience method.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Fatalln package-level convenience method.
func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

// Info package-level convenience method.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof package-level convenience method.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Infoln package-level convenience method.
func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

// Panic package-level convenience method.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf package-level convenience method.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Panicln package-level convenience method.
func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}

// Print package-level convenience method.
func Print(args ...interface{}) {
	defaultLogger.Print(args...)
}

// Printf package-level convenience method.
func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

// Println package-level convenience method.
func Println(args ...interface{}) {
	defaultLogger.Println(args...)
}

// Warn package-level convenience method.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf package-level convenience method.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Warning package-level convenience method.
func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

// Warningf package-level convenience method.
func Warningf(format string, args ...interface{}) {
	defaultLogger.Warningf(format, args...)
}

// Warningln package-level convenience method.
func Warningln(args ...interface{}) {
	defaultLogger.Warningln(args...)
}

// Warnln package-level convenience method.
func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}
