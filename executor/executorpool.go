// MIT License
// Author: Umesh Patil, Neosemantix, Inc.
package executor

// Holds two separate arrays of executors - one for blocking tasks and
// the other for async execution.
// ExecutorPool also fulfills the actual Executor contract:
// Start, Stop, Submit and other methods. That makes it consistent.
type ExecutorPool struct {
	asyncExecutors    []Executor
	blockingExecutors []Executor
}

type ExecPoolCfg struct {

	// Number of executors which will be used to handle async tasks
	AsyncTaskExecutorCount		int

	// Number of executors which will be used to hand blocking task,
	// caller is waiting for the execution result.
	BlockingTaskExecutorCount	int
}

// async: how many executors for execution of async tasks
// blocked: how many executors for execution of blocked tasks
// wfa: wait for availability in the queue for an executor
func NewExecutorPool(async int, blocked int, qc int, wfa bool) *ExecutorPool {
	es := new(ExecutorPool)
	es.asyncExecutors = make([]Executor, async)
	for i := 0; i < async; i++ {
		es.asyncExecutors[i] = NewExecutor(qc, wfa)
	}
	es.blockingExecutors = make([]Executor, blocked)
	for i := 0; i < blocked; i++ {
		es.blockingExecutors[i] = NewExecutor(qc, wfa)
	}
	return es
}

func (es *ExecutorPool) Start() {
	for _, ae := range es.asyncExecutors {
		ae.Start()
	}
	for _, be := range es.blockingExecutors {
		be.Start()
	}
}

func (es *ExecutorPool) Submit(tsk Task) error {
	blocking := tsk.IsBlocking()
	if blocking {
		index := 0
		minEs := es.blockingExecutors[0].HowManyInQueue()
		for i := 1; i < len(es.blockingExecutors); i++ {
			if es.blockingExecutors[i].HowManyInQueue() < minEs {
				index = i
			}
		}
		return es.blockingExecutors[index].Submit(tsk)
	} else {
		index := 0
		minEs := es.asyncExecutors[0].HowManyInQueue()
		for i := 1; i < len(es.asyncExecutors); i++ {
			if es.asyncExecutors[i].HowManyInQueue() < minEs {
				index = i
			}
		}
		return es.asyncExecutors[index].Submit(tsk)
	}
}

func (es *ExecutorPool) HowManyInQueue() int {
	tasksInQueue := 0
	for _, ae := range es.asyncExecutors {
		tasksInQueue += ae.HowManyInQueue()
	}
	for _, be := range es.blockingExecutors {
		tasksInQueue += be.HowManyInQueue()
	}
	return tasksInQueue
}

func (es *ExecutorPool) Stop() {
	for _, ae := range es.asyncExecutors {
		ae.Stop()
	}
	for _, be := range es.blockingExecutors {
		be.Stop()
	}
}

func (es *ExecutorPool) TotalExecutorCount() int {
	return len(es.asyncExecutors) + len(es.blockingExecutors)
}
