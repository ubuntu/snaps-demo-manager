package logger

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

var (
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
)

const (
	normalLogFlags = log.Ldate | log.Ltime | log.Lshortfile
	debugLogFlags  = log.Ldate | log.Ltime
)

func init() {
	debug = log.New(ioutil.Discard, "DEBUG: ", normalLogFlags)
	info = log.New(ioutil.Discard, "INFO: ", normalLogFlags)
	warning = log.New(os.Stderr, "WARNING: ", normalLogFlags)
	err = log.New(os.Stderr, "ERROR: ", normalLogFlags)

	// Note: we need to do this here as some other packages depending on logger during their init() might
	// logs content.
	debug := flag.Bool("debug", false, "Enable debug and info messages")
	info := flag.Bool("info", false, "Enable info messages")
	flag.Parse()

	// enable debug level
	if *info {
		EnableInfo()
		Info("Info message level enabled")
	}
	if *debug {
		EnableDebug()
		Debug("Debug message level enabled")
	}
}

// EnableDebug prints debug messages with all details.
func EnableDebug() {
	EnableInfo()
	debug.SetOutput(os.Stdout)
	debug.SetFlags(debugLogFlags)
	info.SetFlags(debugLogFlags)
	warning.SetFlags(debugLogFlags)
	err.SetFlags(debugLogFlags)
}

// EnableInfo prints info messages without additional details.
func EnableInfo() {
	info.SetOutput(os.Stdout)
}

// NormalLogging returns to warning and err only logging state
func NormalLogging() {
	debug.SetOutput(ioutil.Discard)
	info.SetOutput(ioutil.Discard)
	debug.SetFlags(normalLogFlags)
	info.SetFlags(normalLogFlags)
	warning.SetFlags(normalLogFlags)
	err.SetFlags(normalLogFlags)
}

// Debug formatted message.
func Debug(format string, v ...interface{}) {
	debug.Printf(format, v...)
}

// Info formatted message.
func Info(format string, v ...interface{}) {
	info.Printf(format, v...)
}

// Warning formatted message.
func Warning(format string, v ...interface{}) {
	warning.Printf(format, v...)
}

// Err formatted message.
func Err(format string, v ...interface{}) {
	err.Printf(format, v...)
}
