/*
 * MIT License
 * Copyright (c) 2020. Neosemantix, Inc.
 * Author: Umesh Patil
 */

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

func GetCfgHomeDir() (string, error) {
	err := cleanenv.ReadEnv(&cfgHome)
	return cfgHome.Dir, err
}
