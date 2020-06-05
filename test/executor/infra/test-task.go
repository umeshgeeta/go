// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package infra

import (
	"../../../src/executor"
	"time"
)

const TestTaskExecDurationLowerLimit = 8
const TestTaskExecDurationUpperLimit = 128
const TestTaskDurationRange = TestTaskExecDurationUpperLimit - TestTaskExecDurationLowerLimit

var taskIdCounter int

type TestTask struct {
	id           int
	blocking     bool
	execDuration int
	rc           chan executor.Response
}

func (tt *TestTask) GetId() int {
	return tt.id
}

func (tt *TestTask) Execute() executor.Response {
	resp := executor.NewResponse(tt.id)
	time.Sleep(time.Duration(tt.execDuration) * time.Microsecond)
	// we regard the task is completed successfully, so we set the status
	resp.Status = executor.TaskStatusCompletedSuccessfully
	return *resp
}

func (tt *TestTask) SetRespChan(rc chan executor.Response) {
	tt.rc = rc
}

func (tt *TestTask) GetRespChan() chan executor.Response {
	return tt.rc
}

func (tt *TestTask) IsBlocking() bool {
	return tt.blocking
}

func SetupTestTask() {
	taskIdCounter = 0
}

func NewTestTask(ed int) *TestTask {
	tt := new(TestTask)
	taskIdCounter = +nextTaskId()
	tt.id = taskIdCounter
	tt.blocking = GetRandomBoolean()
	tt.execDuration = ed
	return tt
}

func nextTaskId() int {
	taskIdCounter += 1
	return taskIdCounter
}

func RandomTestTaskExecTime() int {
	return GlobalRand.Intn(TestTaskDurationRange)
}
