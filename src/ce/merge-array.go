// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

// These two arrays contain start and end times for meeting slots. The code should merge these two lists
// and give off a single list that merges overlapping intervals.
// Ex: 4,5 - 5,9 and 8,10 all overlap so they are condensed into 4-10

//Array 1 [1,2] , [5,9], [12,14]
//Array 2 [4,5], [8, 10], [11,12]

//Expected Output[1,2], [4,10], [11,14]

package main

import (
	"fmt"
	"sort"
)

type meetingSlot struct {
	start int
	end   int
}

type meetingSlotSList struct {
	meetings []meetingSlot
}

func (msl *meetingSlotSList) Len() int {
	return len(msl.meetings)
}

func (msl *meetingSlotSList) Less(i, j int) bool {
	result := false
	if msl.meetings[i].start < msl.meetings[j].start {
		result = true
	}
	return result
}

func (msl *meetingSlotSList) Swap(i, j int) {
	temp := msl.meetings[j]
	msl.meetings[j] = msl.meetings[i]
	msl.meetings[i] = temp
}

func mergeMeetingSlots(ar1 *meetingSlotSList, ar2 *meetingSlotSList) *meetingSlotSList {
	var result *meetingSlotSList
	if ar1 == nil {
		result = ar2
	} else if ar2 == nil {
		result = ar1
	} else {
		// both are non-nil
		result = new(meetingSlotSList)
		for _, ms := range ar1.meetings {
			result.meetings = append(result.meetings, ms)
		}
		for _, ms := range ar2.meetings {
			result.meetings = append(result.meetings, ms)
		}
		fmt.Printf("%v\n", result)
		sort.Sort(result)
		fmt.Printf("%v\n", result)
		finalResult := new(meetingSlotSList)
		ms := result.meetings[0]
		runningStart := ms.start
		runningEnd := ms.end
		i := 1
		tuples := len(result.meetings)
		for i < tuples {
			next := result.meetings[i]
			if next.start <= runningEnd {
				// keep consuming slots
				runningEnd = next.end
			} else {
				t := meetingSlot{runningStart, runningEnd}
				finalResult.meetings = append(finalResult.meetings, t)
				runningStart = next.start
				runningEnd = next.end
			}
			i++
		}
		finalResult.meetings = append(finalResult.meetings, meetingSlot{runningStart, runningEnd})
		result = finalResult
	}
	return result
}

func main() {
	ar1 := new(meetingSlotSList)

	ms := meetingSlot{1, 2}
	ar1.meetings = append(ar1.meetings, ms)
	ms = meetingSlot{12, 14}
	ar1.meetings = append(ar1.meetings, ms)
	ms = meetingSlot{5, 9}
	ar1.meetings = append(ar1.meetings, ms)

	ar2 := new(meetingSlotSList)

	ms = meetingSlot{11, 12}
	ar2.meetings = append(ar2.meetings, ms)
	ms = meetingSlot{4, 5}
	ar2.meetings = append(ar2.meetings, ms)
	ms = meetingSlot{8, 10}
	ar2.meetings = append(ar2.meetings, ms)

	result := mergeMeetingSlots(ar1, ar2)

	fmt.Printf("%v\n", result)
}
