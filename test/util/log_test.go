// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"../../util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLog(t *testing.T) {
	assert := assert.New(t)
	assert.False(util.IsLoggingConfigured())

	util.InitializeLog("./log/test.log", 10, 2, 5, false)
	util.SetConsoleLog(true)
	util.Log("Started log")
	util.SetDeubgLog(true)
	util.LogDebug("Debug log should come up as well")
	util.Log("End log")

	assert.True(util.IsLoggingConfigured())
}
