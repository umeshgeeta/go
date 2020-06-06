// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	logutil "../../src/util"
)

// Simple test to ensure that logs are created correctly.
func main() {
	logutil.InitializeLog("./log/test.log", 10, 2, 5, false)
	logutil.GlobalLogSettings.LogOnConsole = true
	logutil.Log("Started log")
	logutil.GlobalLogSettings.DebugLog = true
	logutil.LogDebug("Debug log should come up as well")
	logutil.Log("End log")
}
