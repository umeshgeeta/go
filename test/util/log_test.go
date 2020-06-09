// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"../../util"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	assert := assert.New(t)
	assert.False(util.IsLoggingConfigured())

	logFileName := "./log/test.log"

	util.InitializeLog(logFileName, 10, 2, 5, false)

	if _, err := os.Stat(logFileName); os.IsNotExist(err) {
		t.Errorf("Test Failed error: %v\n", err)
	}

	util.SetConsoleLog(true)
	util.Log("Started log")
	util.SetDeubgLog(true)
	util.LogDebug("Debug log should come up as well")
	util.Log("End log")

	assert.True(util.IsLoggingConfigured())
}
