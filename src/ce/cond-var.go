// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wt1 := new(waitingTask)
	wt1.wtc = sync.NewCond(wt1)

	go func(wt *waitingTask) {
		wt.Lock()
		fmt.Println("Waiting started at: " + time.Now().String())
		for len(wt.response) == 0 {
			wt.wtc.Wait()
		}
		fmt.Println(fmt.Sprintf("Waiting done, response is %s\n", wt.response))
		wt.Unlock()
	}(wt1)

	time.Sleep(1 * time.Second)
	wt1.response = "Come out: " + time.Now().String()
	wt1.wtc.Signal()
	time.Sleep(1 * time.Second)
	fmt.Println("Everything done at: " + time.Now().String())
}

type waitingTask struct {
	sync.Mutex
	response string
	wtc      *sync.Cond
}
