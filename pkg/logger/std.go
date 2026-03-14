package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// StdLogger 标准输出日志记录�?
type StdLogger struct {
	logger *log.Logger
}

// NewStdLogger 创建标准输出日志记录�?
func NewStdLogger() Logger {
	return &StdLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Debug 调试日志
func (l *StdLogger) Debug(msg string, fields ...Field) {
	l.log("DEBUG", msg, fields...)
}

// Info 信息日志
func (l *StdLogger) Info(msg string, fields ...Field) {
	l.log("INFO", msg, fields...)
}

// Warn 警告日志
func (l *StdLogger) Warn(msg string, fields ...Field) {
	l.log("WARN", msg, fields...)
}

// Error 错误日志
func (l *StdLogger) Error(msg string, fields ...Field) {
	l.log("ERROR", msg, fields...)
}

// log 内部日志输出方法
func (l *StdLogger) log(level, msg string, fields ...Field) {
	var parts []string
	parts = append(parts, fmt.Sprintf("[%s] %s", level, msg))

	if len(fields) > 0 {
		var fieldStrs []string
		for _, field := range fields {
			fieldStrs = append(fieldStrs, fmt.Sprintf("%s=%v", field.Key, field.Value))
		}
		parts = append(parts, strings.Join(fieldStrs, " "))
	}

	l.logger.Println(strings.Join(parts, " | "))
}

