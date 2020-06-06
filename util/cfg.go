// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Package util contains various types and functions used in ancillary mode.
package util

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type ConfigHome struct {
	Dir string `env:"GO_CFG_HOME" env-description:"Directory where we can find s3downloader.cfg file"`
}

var cfgHome ConfigHome

// Read configurations
func ReadCfg(cfg interface{}, fileName string) error {
	err := cleanenv.ReadEnv(&cfgHome)
	if err != nil {
		log.Fatal("Error reading environment variable GO_CFG_HOME")
	}
	if len(cfgHome.Dir) == 0 {
		log.Fatal("Environment variable GO_CFG_HOME is not defined")
	}
	cfgFileName := cfgHome.Dir + "/" + fileName
	fmt.Printf("Config file in use %s\n", cfgFileName)
	return cleanenv.ReadConfig(cfgFileName, cfg)
}

// Returns directory which holds config files.
func GetCfgHomeDir() (string, error) {
	err := cleanenv.ReadEnv(&cfgHome)
	return cfgHome.Dir, err
}
