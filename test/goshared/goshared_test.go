// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// Contains test programs to check importing of packages from 'goshared'.
package goshared

import (
	"fmt"
	"github.com/umeshgeeta/goshared/executor"
	"github.com/umeshgeeta/goshared/util"
	"testing"
)

func TestExecutorServiceFromDefaultCfg(t *testing.T) {
	es := executor.NewExecutionService("xxx", true)
	util.Log(fmt.Sprintf("Config: %v\n", executor.GlobalExecServiceCfg))
	util.Log(fmt.Sprintf("%v\n", es))
	fmt.Println("Done goshared_test")
}
