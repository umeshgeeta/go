// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package executor

const DefaultTaskQueueCapacity = 5
const WaitForExecutorAvailDefault = false
const WaitForChannelAvailDefault = false
const DefaultTaskResultChannelCapacity = 1

type ExecutionService struct {
	TaskDispatcher *Dispatcher
}

func NewExecutionService(execsForBlockingTasks int, execsForAsyncTasks int) *ExecutionService {
	es := new(ExecutionService)
	es.TaskDispatcher = NewDispatcher(execsForBlockingTasks+execsForAsyncTasks,
		DefaultTaskResultChannelCapacity,
		NewExecutorPool(execsForAsyncTasks,
			execsForBlockingTasks,
			DefaultTaskQueueCapacity,
			WaitForExecutorAvailDefault),
		WaitForChannelAvailDefault)
	return es
}

func (es *ExecutionService) Start() {
	es.TaskDispatcher.Start()
}

func (es *ExecutionService) Submit(tsk Task) (error, *Response) {
	return es.TaskDispatcher.Submit(tsk)
}

func (es *ExecutionService) Stop() {
	es.TaskDispatcher.Stop()
}
