// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"../../executor"
	"./infra"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	infra.SetupRand()
	infra.SetupTestTask()

	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}

func TestExecutorFailSubmission(t *testing.T) {
	//t.SkipNow()
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

	var wg sync.WaitGroup
	wg.Add(2)

	if err == nil {
		go waitForResponse(ch, t, executor.TaskStatusCompletedSuccessfully, "Tasks Status as expected.", &wg)
	} else {
		t.Errorf("Task submission error: %v\n", err)
	}

	task2 := infra.NewTestTask(10)
	ch2 := make(chan executor.Response)
	task2.SetRespChan(ch2)
	err2 := thread.Submit(task2)

	if err2 == nil {
		fmt.Printf("In queue: %d\n", thread.HowManyInQueue())
		go waitForResponse(ch2, t, executor.TaskStatusCompletedSuccessfully, "Tasks Status as expected.", &wg)
	} else {
		t.Errorf("Task submission error: %v\n", err)
	}

	wg.Wait()
	fmt.Println("wait wg complete")
	thread.Stop()
	fmt.Println("thread stopped, everything is done")
}

func waitForResponse(ch chan executor.Response, t *testing.T, ts int, msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("waitForResponse Start")
	rsp := <-ch
	assert := assert.New(t)
	assert.Equal(rsp.Status, ts, msg)
	fmt.Println("waitForResponse Done")
}
