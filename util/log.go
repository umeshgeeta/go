// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// There are three methods to initialize logs and set log settings:
//
// 1) Caller can pass LoggingCfg structure as defined here to initialize logs.
//
// 2) Alternatively, assuming caller has used 'cfg' program in the 'util' package,
// that 'uber' configuration file can be passed as long as it has a JSON member
// of name "LogSettings". Note the environment variable GO_CFG_HOME needs to be
// set for this option to work. The value of this environmental variable should
// point to a directory which contains the configuration file which has the
// JSON member LogSettings.
//
// 3) Finally, caller can explicitly call InitializeLog method with arguments.
package util

import (
	"encoding/json"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"path/filepath"
)

type LoggingCfg struct {
	LogFileName  string
	MaxSizeInMb  int
	Backups      int
	AgeInDays    int
	Compress     bool
	LogOnConsole bool
	DebugLog     bool
}

var GlobalLogSettings *LoggingCfg = new(LoggingCfg) // all bool settings false by default

// Initialize logging to given inputs:
// fn	:	Log files with fill path
// ms	:	Maximum allowed log file size  in Megabytes
// bk	:	How many backups to be retained.
// age	:	Past logs of how many days to be retained.
// compress:	Whether logs are compressed or not.
func InitializeLog(fn string, ms int, bk int, age int, compress bool) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   fn,
		MaxSize:    ms, // megabytes
		MaxBackups: bk,
		MaxAge:     age,      //days
		Compress:   compress, // disabled by default
	})
	logFilePath, _ := filepath.Abs(fn)
	log.Printf("logFilePath: %s\n", logFilePath)
	if GlobalLogSettings.LogOnConsole {
		fmt.Printf("logFilePath: %s\n", logFilePath)
	}
}

func SetLoggingCfg(ls *LoggingCfg) {
	if ls != nil {
		GlobalLogSettings = ls
		InitializeLog(ls.LogFileName, ls.MaxSizeInMb, ls.Backups, ls.AgeInDays, ls.Compress)
	} else {
		log.Fatal("Logging configuration is nil")
	}
}

// It assumes argument configuration file in JSON format and it contains an
// element named LogSettings. Contents of that member are used to build the
// log setting configuration.
func SetLogSettings(cfgFileName string) {
	if len(cfgFileName) > 0 {
		ls, err := extractLogSettings(cfgFileName)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error extracting LogSettings from the given config file (%s): %v\n", cfgFileName, err))
		} else {
			ls, err := formLoggingCfg(ls)
			if err != nil {
				fmt.Println(fmt.Sprintf("Error forming LoggingCfg from the given config file (%s): %v\n", cfgFileName, err))
			} else {
				SetLoggingCfg(ls)
				// Log the setting values which will be used
				Log(fmt.Sprintf("Log Settings: %v\n", ls))
			}
		}
	} else {
		msg := "configuration file is nil"
		fmt.Println(msg)
		log.Fatal(msg)
	}
}

func formLoggingCfg(ls string) (*LoggingCfg, error) {
	var lc LoggingCfg
	err := json.Unmarshal([]byte(ls), &lc)
	return &lc, err
}

// Enable or disable the console logging
func SetConsoleLog(val bool) {
	GlobalLogSettings.LogOnConsole = val
}

// Enable or disable debug logging.
func SetDeubgLog(val bool) {
	GlobalLogSettings.DebugLog = val
}

// Log the given message.
func Log(msg string) {
	log.Println(msg)
	if GlobalLogSettings.LogOnConsole {
		fmt.Println(msg)
	}
}

// Log debug messages. Invocation of this call will result in adding the message
// to the log provided SetDebugLog(true) is called.
func LogDebug(msg string) {
	if GlobalLogSettings.DebugLog {
		Log(msg)
	}
}
