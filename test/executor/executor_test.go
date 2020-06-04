/*
 * MIT License
 * Author: Umesh Patil, Neosemantix, Inc.
 */
package main

import (
	"../../src/executor"
	"./infra"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	infra.SetupRand()
	infra.SetupTestTask()

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestExecutorFailSubmission(t *testing.T) {
	thread := executor.NewExecutor(2, false)
	task := infra.NewTestTask(infra.RandomTestTaskExecTime())
	err := thread.Submit(task)
	assert := assert.New(t)
	assert.NotNil(err, "Expected error: %v\n", err)
}

func TestExecutorExecutionSuccess(t *testing.T) {
	thread := executor.NewExecutor(2, false)
	thread.Start()
	task := infra.NewTestTask(10000)
	ch := make(chan executor.Response)
	task.SetRespChan(ch)
	err := thread.Submit(task)
	if err == nil {
		rsp := <-ch
		assert := assert.New(t)
		assert.Equal(rsp.Status, executor.TaskStatusCompletedSuccessfully, "Tasks Status as expected.")
	} else {
		t.Errorf("Task submission error: %v\n", err)
	}
	task2 := infra.NewTestTask(10)
	ch2 := make(chan executor.Response)
	task2.SetRespChan(ch2)
	err2 := thread.Submit(task2)
	if err2 == nil {
		fmt.Printf("In queue: %d\n", thread.HowManyInQueue())
		rsp := <-ch2
		assert := assert.New(t)
		assert.Equal(rsp.Status, executor.TaskStatusCompletedSuccessfully, "Tasks Status as expected.")
	} else {
		t.Errorf("Task submission error: %v\n", err)
	}
	thread.Stop()
}
