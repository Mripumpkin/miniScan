package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	config "miniScan/utils/conf"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

// Define log instances
var (
	defaultLogger *logrus.Logger
	LogPortScan   *logrus.Logger
	LogDomain     *logrus.Logger
	LogWebInfo    *logrus.Logger
	once          sync.Once
)

func init() {
	// Initialize log instances with singleton pattern
	once.Do(func() {
		defaultLogger = newLogrusLogger(config.Config())
		LogPortScan = newLogrusLoggerWithRotation(config.Config(), "portscan", "Port Scan")
		LogDomain = newLogrusLoggerWithRotation(config.Config(), "domain", "Domain Scan")
		LogWebInfo = newLogrusLoggerWithRotation(config.Config(), "web", "Web Probe")
	})
}

// NewLogger returns a configured logrus instance
func NewLogger(cfg config.Provider) *logrus.Logger {
	once.Do(func() {
		defaultLogger = newLogrusLogger(cfg)
	})
	return defaultLogger
}

// newLogrusLogger creates a basic logrus instance
func newLogrusLogger(cfg config.Provider) *logrus.Logger {
	l := logrus.New()

	if cfg.GetBool("json_logs") {
		l.Formatter = new(logrus.JSONFormatter)
	} else {
		// Use custom TextFormatter
		l.Formatter = &CustomTextFormatter{Prefix: "LOG"}
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

// newLogrusLoggerWithRotation creates a logrus instance with log rotation and also writes to console
func newLogrusLoggerWithRotation(cfg config.Provider, filename string, prefix string) *logrus.Logger {
	l := logrus.New()

	// Set filename with function name and date
	date := time.Now().Format("2006-01-02")
	logFile := filepath.Join("logs", filename+"-"+date+".log")

	// Use lumberjack to handle log size and rotation
	lumberjackLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1,    // Maximum file size of 1MB
		MaxBackups: 10,   // Keep the last 10 backups
		MaxAge:     30,   // Keep logs for a maximum of 30 days
		Compress:   true, // Compress old log files
	}

	// Create multi-writer for both stdout and file
	mw := io.MultiWriter(os.Stdout, lumberjackLogger)
	l.Out = mw

	// Use custom TextFormatter
	l.Formatter = &CustomTextFormatter{Prefix: prefix}

	// Set log level based on configuration
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

	return l
}

// CustomTextFormatter defines a custom text formatter
type CustomTextFormatter struct {
	Prefix string
}

// Format formats the log output
func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Get caller information
	_, file, line, ok := runtime.Caller(6)
	if !ok {
		file = "unknown"
		line = 0
	}

	// Format timestamp
	timestamp := entry.Time.Format("2006-01-02 15:04:05,999")

	// Format log level
	level := entry.Level.String()

	// Set color for log level
	var levelColor, resetColor string
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = "\033[34m" // Blue
	case logrus.InfoLevel:
		levelColor = "\033[34m" // Green
	case logrus.WarnLevel:
		levelColor = "\033[33m" // Yellow
	case logrus.ErrorLevel:
		levelColor = "\033[31m" // Red
	case logrus.FatalLevel, logrus.PanicLevel:
		levelColor = "\033[35m" // Magenta
	default:
		levelColor = "\033[0m" // Default
	}
	resetColor = "\033[0m"

	// Format message
	msg := entry.Message

	// Generate log line
	lineText := fmt.Sprintf("[%s%s%s][%s%s%s][%s] %s (%s:%d)\n",
		levelColor, f.Prefix, resetColor,
		levelColor, level, resetColor,
		timestamp,
		msg,
		filepath.Base(file),
		line,
	)

	return []byte(lineText), nil
}

// Fields is a map string interface to define fields in the structured log
type Fields map[string]interface{}

// With allows us to define fields in our structured logs
func (f Fields) With(k string, v interface{}) Fields {
	f[k] = v
	return f
}

// WithFields allows us to define fields in our structured logs
func (f Fields) WithFields(f2 Fields) Fields {
	for k, v := range f2 {
		f[k] = v
	}
	return f
}

// Debug is a package-level convenience method.
func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

// Debugf is a package-level convenience method.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Debugln is a package-level convenience method.
func Debugln(args ...interface{}) {
	defaultLogger.Debugln(args...)
}

// Error is a package-level convenience method.
func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

// Errorf is a package-level convenience method.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Errorln is a package-level convenience method.
func Errorln(args ...interface{}) {
	defaultLogger.Errorln(args...)
}

// Fatal is a package-level convenience method.
func Fatal(args ...interface{}) {
	defaultLogger.Fatal(args...)
}

// Fatalf is a package-level convenience method.
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// Fatalln is a package-level convenience method.
func Fatalln(args ...interface{}) {
	defaultLogger.Fatalln(args...)
}

// Info is a package-level convenience method.
func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

// Infof is a package-level convenience method.
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Infoln is a package-level convenience method.
func Infoln(args ...interface{}) {
	defaultLogger.Infoln(args...)
}

// Panic is a package-level convenience method.
func Panic(args ...interface{}) {
	defaultLogger.Panic(args...)
}

// Panicf is a package-level convenience method.
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Panicln is a package-level convenience method.
func Panicln(args ...interface{}) {
	defaultLogger.Panicln(args...)
}

// Print is a package-level convenience method.
func Print(args ...interface{}) {
	defaultLogger.Print(args...)
}

// Printf is a package-level convenience method.
func Printf(format string, args ...interface{}) {
	defaultLogger.Printf(format, args...)
}

// Println is a package-level convenience method.
func Println(args ...interface{}) {
	defaultLogger.Println(args...)
}

// Warn is a package-level convenience method.
func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

// Warnf is a package-level convenience method.
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Warning is a package-level convenience method.
func Warning(args ...interface{}) {
	defaultLogger.Warning(args...)
}

// Warningf is a package-level convenience method.
func Warningf(format string, args ...interface{}) {
	defaultLogger.Warningf(format, args...)
}

// Warningln is a package-level convenience method.
func Warningln(args ...interface{}) {
	defaultLogger.Warningln(args...)
}

// Warnln is a package-level convenience method.
func Warnln(args ...interface{}) {
	defaultLogger.Warnln(args...)
}
