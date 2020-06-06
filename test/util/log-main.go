// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"github.com/umeshgeeta/go/util"
)

// Simple test to ensure that logs are created correctly.
func main() {
	util.InitializeLog("./log/test.log", 10, 2, 5, false)
	util.GlobalLogSettings.LogOnConsole = true
	util.Log("Started log")
	util.GlobalLogSettings.DebugLog = true
	util.LogDebug("Debug log should come up as well")
	util.Log("End log")
}
