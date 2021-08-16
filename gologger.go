package gologger

import (
	"io"
	"os"
	"fmt"
	"time"
	"sync"
	"strings"
)
//https://github.com/subchen/go-log
var (
	_ StdLog       = Default
	_ LogInterface = Default
)

// StdLog is interface for builtin log
type StdLog interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// LogInterface is interface for this logger
type LogInterface interface {
	Debug(...interface{})
	Info(...interface{})
	Print(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Panic(...interface{})
	Fatal(...interface{})

	Debugln(...interface{})
	Infoln(...interface{})
	Println(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
	Panicln(...interface{})
	Fatalln(...interface{})

	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Printf(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}

// Format is a interface used to implement a custom Format
type Format interface {
	Format(level Level, msg string, logger *Logger) []byte
}

// simpleFormat is default formmatter
type simpleFormat struct {
}

// Format implements log.Format
func (f *simpleFormat) Format(level Level, msg string, logger *Logger) []byte {
	time := time.Now().Format("15:04:05.000")
	return []byte(fmt.Sprintf("%s %s %s\n", time, level.String(), msg))
}

// Level type
type Level uint32

// These are the different logging levels
const (
	OFF Level = iota
	FATAL
	PANIC
	ERROR
	WARN
	INFO
	DEBUG
)

// String converts the Level to a string
func (level Level) String() string {
	switch level {
	case OFF:
		return "OFF"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	case ERROR:
		return "ERROR"
	case WARN:
		return "WARN"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

// ColorString converts the Level to a string with term colorful
func (level Level) ColorString() string {
	switch level {
	case OFF:
		return "OFF"
	case FATAL:
		return "\033[35mFATAL\033[0m"
	case PANIC:
		return "\033[35mPANIC\033[0m"
	case ERROR:
		return "\033[31mERROR\033[0m"
	case WARN:
		return "\033[33mWARN\033[0m"
	case INFO:
		return "\033[32mINFO\033[0m"
	case DEBUG:
		return "\033[34mDEBUG\033[0m"
	default:
		return "UNKNOWN"
	}
}

// ParseLevel takes a string level and returns the log level constant.
func ParseLevel(name string) (Level, error) {
	switch strings.ToUpper(name) {
	case "OFF":
		return OFF, nil
	case "FATAL":
		return FATAL, nil
	case "PANIC":
		return PANIC, nil
	case "ERROR":
		return ERROR, nil
	case "WARN":
		return WARN, nil
	case "INFO":
		return INFO, nil
	case "DEBUG":
		return DEBUG, nil
	}

	return 0, fmt.Errorf("invalid log.Level: %q", name)
}

// Default is a default Logger instance
var Default = New()

// IsDebugEnabled indicates whether output message
func IsDebugEnabled() bool {
	return Default.IsDebugEnabled()
}

// IsInfoEnabled indicates whether output message
func IsInfoEnabled() bool {
	return Default.IsInfoEnabled()
}

// IsPrintEnabled indicates whether output message
func IsPrintEnabled() bool {
	return Default.IsPrintEnabled()
}

// IsWarnEnabled indicates whether output message
func IsWarnEnabled() bool {
	return Default.IsWarnEnabled()
}

// IsErrorEnabled indicates whether output message
func IsErrorEnabled() bool {
	return Default.IsErrorEnabled()
}

// IsPanicEnabled indicates whether output message
func IsPanicEnabled() bool {
	return Default.IsPanicEnabled()
}

// IsFatalEnabled indicates whether output message
func IsFatalEnabled() bool {
	return Default.IsFatalEnabled()
}

// IsDisabled indicates whether output message
func IsDisabled() bool {
	return Default.IsDisabled()
}

// Debug outputs message, Arguments are handled by fmt.Sprint
func Debug(obj ...interface{}) {
	Default.Debug(obj...)
}

// Info outputs message, Arguments are handled by fmt.Sprint
func Info(obj ...interface{}) {
	Default.Info(obj...)
}

// Print outputs message, Arguments are handled by fmt.Sprint
func Print(obj ...interface{}) {
	Default.Print(obj...)
}

// Warn outputs message, Arguments are handled by fmt.Sprint
func Warn(obj ...interface{}) {
	Default.Warn(obj...)
}

// Error outputs message, Arguments are handled by fmt.Sprint
func Error(obj ...interface{}) {
	Default.Error(obj...)
}

// Panic outputs message, and followed by a call to panic() Arguments are handled by fmt.Sprint
func Panic(obj ...interface{}) {
	Default.Panic(obj...)
}

// Fatal outputs message, and followed by a call to os.Exit(1) Arguments are handled by fmt.Sprint
func Fatal(obj ...interface{}) {
	Default.Fatal(obj...)
}

// Debugln outputs message, Arguments are handled by fmt.Sprintln
func Debugln(obj ...interface{}) {
	Default.Debugln(obj...)
}

// Infoln outputs message, Arguments are handled by fmt.Sprintln
func Infoln(obj ...interface{}) {
	Default.Infoln(obj...)
}

// Println outputs message, Arguments are handled by fmt.Sprintln
func Println(obj ...interface{}) {
	Default.Println(obj...)
}

// Warnln outputs message, Arguments are handled by fmt.Sprintln
func Warnln(obj ...interface{}) {
	Default.Warnln(obj...)
}

// Errorln outputs message, Arguments are handled by fmt.Sprintln
func Errorln(obj ...interface{}) {
	Default.Errorln(obj...)
}

// Panicln outputs message and followed by a call to panic(), Arguments are handled by fmt.Sprintln
func Panicln(obj ...interface{}) {
	Default.Panicln(obj...)
}

// Fatalln outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprintln
func Fatalln(obj ...interface{}) {
	Default.Fatalln(obj...)
}

// Debugf outputs message, Arguments are handled by fmt.Sprintf
func Debugf(msg string, args ...interface{}) {
	Default.Debugf(msg, args...)
}

// Infof outputs message, Arguments are handled by fmt.Sprintf
func Infof(msg string, args ...interface{}) {
	Default.Infof(msg, args...)
}

// Printf outputs message, Arguments are handled by fmt.Sprintf
func Printf(msg string, args ...interface{}) {
	Default.Printf(msg, args...)
}

// Warnf outputs message, Arguments are handled by fmt.Sprintf
func Warnf(msg string, args ...interface{}) {
	Default.Warnf(msg, args...)
}

// Errorf outputs message, Arguments are handled by fmt.Sprintf
func Errorf(msg string, args ...interface{}) {
	Default.Errorf(msg, args...)
}

// Panicf outputs message and followed by a call to panic(), Arguments are handled by fmt.Sprintf
func Panicf(msg string, args ...interface{}) {
	Default.Panicf(msg, args...)
}

// Fatalf outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprintf
func Fatalf(msg string, args ...interface{}) {
	Default.Fatalf(msg, args...)
}

// Exit is equals os.Exit
var Exit = os.Exit

// Logger is represents an active logging object
type Logger struct {
	mutex     sync.Mutex
	Level     Level
	Format 	  Format
	Output    io.Writer
}

// New creates a new Logger
func New() *Logger {
	return &Logger{
		Level:     INFO,
		Format:    new(simpleFormat),
		Output:       os.Stdout,
	}
}

// IsDebugEnabled indicates whether output message
func (l *Logger) IsDebugEnabled() bool {
	return l.Level >= DEBUG
}

// IsInfoEnabled indicates whether output message
func (l *Logger) IsInfoEnabled() bool {
	return l.Level >= INFO
}

// IsPrintEnabled indicates whether output message
func (l *Logger) IsPrintEnabled() bool {
	return l.Level > OFF
}

// IsWarnEnabled indicates whether output message
func (l *Logger) IsWarnEnabled() bool {
	return l.Level >= WARN
}

// IsErrorEnabled indicates whether output message
func (l *Logger) IsErrorEnabled() bool {
	return l.Level >= ERROR
}

// IsPanicEnabled indicates whether output message
func (l *Logger) IsPanicEnabled() bool {
	return l.Level >= PANIC
}

// IsFatalEnabled indicates whether output message
func (l *Logger) IsFatalEnabled() bool {
	return l.Level >= FATAL
}

// IsDisabled indicates whether output message
func (l *Logger) IsDisabled() bool {
	return l.Level <= OFF
}

// Debug outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Debug(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprint(obj...))
	}
}

// Info outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Info(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

// Print outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Print(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprint(obj...))
	}
}

// Warn outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Warn(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprint(obj...))
	}
}

// Error outputs message, Arguments are handled by fmt.Sprint
func (l *Logger) Error(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprint(obj...))
	}
}

// Panic outputs message, and followed by a call to panic() Arguments are handled by fmt.Sprint
func (l *Logger) Panic(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprint(obj...))
	}
	panic(fmt.Sprint(obj...))
}

// Fatal outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprint
func (l *Logger) Fatal(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprint(obj...))
	}
	Exit(1)
}

// Debugln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Debugln(obj ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, vsprintln(obj...))
	}
}

// Infoln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Infoln(obj ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, vsprintln(obj...))
	}
}

// Println outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Println(obj ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, vsprintln(obj...))
	}
}

// Warnln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Warnln(obj ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, vsprintln(obj...))
	}
}

// Errorln outputs message, Arguments are handled by fmt.Sprintln
func (l *Logger) Errorln(obj ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, vsprintln(obj...))
	}
}

// Panicln outputs message and followed by a call to panic(), Arguments are handled by fmt.Sprintln
func (l *Logger) Panicln(obj ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, vsprintln(obj...))
	}
	panic(vsprintln(obj...))
}

// Fatalln outputs message and followed by a call to os.Exit(1), Arguments are handled by fmt.Sprintln
func (l *Logger) Fatalln(obj ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, vsprintln(obj...))
	}
	Exit(1)
}

// Debugf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Debugf(msg string, args ...interface{}) {
	if l.Level >= DEBUG {
		l.log(DEBUG, fmt.Sprintf(msg, args...))
	}
}

// Infof outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Infof(msg string, args ...interface{}) {
	if l.Level >= INFO {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

// Printf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Printf(msg string, args ...interface{}) {
	if l.Level != OFF {
		l.log(INFO, fmt.Sprintf(msg, args...))
	}
}

// Warnf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Warnf(msg string, args ...interface{}) {
	if l.Level >= WARN {
		l.log(WARN, fmt.Sprintf(msg, args...))
	}
}

// Errorf outputs message, Arguments are handles by fmt.Sprintf
func (l *Logger) Errorf(msg string, args ...interface{}) {
	if l.Level >= ERROR {
		l.log(ERROR, fmt.Sprintf(msg, args...))
	}
}

// Panicf outputs message and followed by a call to panic(), Arguments are handles by fmt.Sprintf
func (l *Logger) Panicf(msg string, args ...interface{}) {
	if l.Level >= PANIC {
		l.log(PANIC, fmt.Sprintf(msg, args...))
	}
	panic(fmt.Sprintf(msg, args...))
}

// Fatalf outputs message and followed by a call to os.Exit(1), Arguments are handles by fmt.Sprintf
func (l *Logger) Fatalf(msg string, args ...interface{}) {
	if l.Level >= FATAL {
		l.log(FATAL, fmt.Sprintf(msg, args...))
	}
	Exit(1)
}

func (l *Logger) log(level Level, msg string) {
	line := l.Format.Format(level, msg, l)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	_, err := l.Output.Write(line)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write log, %v\n", err)
	}
}

// vsprintln => spaces are always added between operands
func vsprintln(obj ...interface{}) string {
	msg := fmt.Sprintln(obj...)
	return msg[:len(msg)-1]
}