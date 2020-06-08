// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Package util contains common functionality which is required by a typical
// server side any GO application. Namely it contains functionality to read
// JSON configuration file (from ilyakaznacheev) and rotating logging
// functionality (from lumberjack). Essentially these are simply few useful
// wrappers convenient for any GO application.
//
// Environment variable GO_CFG_HOME is checked to see if contains provided
// configuration file. If found, it is read and configuration structure is
// returned. If this environmental variable is not set, no configuration is
// read or returned.
package util

import (
	"encoding/json"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"io/ioutil"
	"log"
	"os"
)

type ConfigHome struct {
	Dir string `env:"GO_CFG_HOME" env-description:"Directory where we can find configuration file"`
}

var cfgHome ConfigHome

// Read configurations
func ReadCfg(cfg interface{}, fileName string) error {
	err := cleanenv.ReadEnv(&cfgHome)
	if err != nil {
		msg := fmt.Sprintf("Error reading environment variable GO_CFG_HOME: %v", err)
		fmt.Println(msg)
		log.Fatal(msg)
	}
	if len(cfgHome.Dir) == 0 {
		msg := fmt.Sprintf("Environment variable GO_CFG_HOME is not defined")
		fmt.Println(msg)
		log.Fatal(msg)
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

// Return the byte array of specified config element in the passed filename.
func ExtractCfgJsonElement(fileName string, cfgJsonElementName string) (string, error) {
	err := cleanenv.ReadEnv(&cfgHome)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error reading environment variable GO_CFG_HOME: %v", err))
	}
	if len(cfgHome.Dir) == 0 {
		fmt.Println(fmt.Sprintf("Environment variable GO_CFG_HOME is not defined"))
	}
	cfgFileName := cfgHome.Dir + "/" + fileName
	// Open our jsonFile
	jsonFile, err := os.Open(cfgFileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		return "", err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	// get value of requested cfg element
	rr := result[cfgJsonElementName]
	fmt.Printf("LogSettings: %v\n", rr)
	buf, err := json.Marshal(rr)
	fmt.Println(string(buf))
	return string(buf), err
}
