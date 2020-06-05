// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package util

import (
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"path/filepath"
)

type LogSettings struct {
	LogOnConsole	bool
	DebugLog		bool
}

var GlobalLogSettings *LogSettings = new(LogSettings)	// all bool settings false by default

func InitializeLog(fn string, ms int, bk int, age int, compress bool) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   fn,
		MaxSize:    ms, // megabytes
		MaxBackups: bk,
		MaxAge:     age, //days
		Compress:   compress,     // disabled by default
	})
	logFilePath, _ := filepath.Abs(fn)
	log.Printf("logFilePath: %s\n", logFilePath)
	if GlobalLogSettings.LogOnConsole {
		fmt.Printf("logFilePath: %s\n", logFilePath)
	}
}

func Log(msg string) {
	log.Println(msg)
	if GlobalLogSettings.LogOnConsole {
		fmt.Println(msg)
	}
}

func LogDebug(msg string) {
	if GlobalLogSettings.DebugLog {
		Log(msg)
	}
}
