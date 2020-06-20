// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"container/list"
	"errors"
	"fmt"
)

type tuple struct {
	smallerChar rune
	biggerChar  rune
}

func (t tuple) string() string {
	return "(" + string(t.smallerChar) + ", " + string(t.biggerChar) + ")"
}

type orderedChars struct {
	charList *list.List
}

func newOc() *orderedChars {
	result := new(orderedChars)
	result.charList = list.New()
	return result
}

func newOcForTuple(t tuple) *orderedChars {
	result := newOc()
	first := result.charList.PushFront(t.smallerChar)
	result.charList.InsertAfter(t.biggerChar, first)
	return result
}

func (oc *orderedChars) strRep() string {
	result := "["
	mark := oc.charList.Front()
	for mark != nil {
		result = result + string(mark.Value.(rune))
		mark = mark.Next()
	}
	result = result + "]"
	return result
}

func (oc *orderedChars) firstChar() rune {
	return oc.charList.Front().Value.(rune)
}

func (oc *orderedChars) lastChar() rune {
	return oc.charList.Back().Value.(rune)
}

func (oc *orderedChars) isEmpty() bool {
	if oc.charList.Len() > 0 {
		return true
	} else {
		return false
	}
}

func (oc *orderedChars) insertAtStart(v interface{}) {
	start := oc.charList.Front()
	oc.charList.InsertBefore(v, start)
}

func (oc *orderedChars) appendElement(v interface{}) {
	end := oc.charList.Back()
	oc.charList.InsertAfter(v, end)
}

// Found value, second returned value:
//					-1:	contradiction
//					0:	no smaller or bigger char of tuple was found
//					1:	only smaller of tuple was found
//					2:	only bigger of tuple was found
//					3:	both smaller and bigger were found
//					11:	Smaller appended at the start
//					12: Bigger appended at the start
func (oc *orderedChars) consume(t tuple) (*orderedChars, int, error) {
	found := -1
	if t.biggerChar == oc.firstChar() {
		oc.insertAtStart(t.smallerChar)
		found = 11
		return nil, found, nil
	} else if t.smallerChar == oc.lastChar() {
		oc.appendElement(t.biggerChar)
		found = 12
		return nil, found, nil
	} else {
		return oc.clone(t)
	}
}

func (oc *orderedChars) clone(t tuple) (*orderedChars, int, error) {
	found := -1
	var result *orderedChars = nil
	var err error
	mark := oc.charList.Front()
	var sc, bc *list.Element = nil, nil
	for mark != nil {
		if mark.Value.(rune) == t.smallerChar {
			if bc != nil {
				// we got bigger character before the smaller character
				err = errors.New(fmt.Sprintf("Contrdiction for tuple %v", t))
			}
			sc = mark
		} else if mark.Value.(rune) == t.biggerChar {
			bc = mark
		}
		mark = mark.Next()
	}
	if sc != nil {
		if bc != nil {
			// we found both chars of tuple in the OC, nothing to do
			found = 3
		} else {
			result = oc.cloneSlice(oc.charList.Front(), sc)
			result.appendElement(bc)
			found = 1
		}
	} else {
		if bc != nil {
			result = oc.cloneSlice(bc, oc.charList.Back())
			result.insertAtStart(sc.Value)
			found = 2
		} else {
			// both sc and bc were not found, we return
			found = 0
		}
	}
	return result, found, err
}

func (oc orderedChars) cloneSlice(start *list.Element, end *list.Element) *orderedChars {
	result := newOc()
	mark := start
	for mark != end {
		result.appendElement(mark.Value)
		mark = mark.Next()
	}
	// last element is not added, we want to have it added as well
	result.charList.PushBack(mark.Value)
	return result
}

func (oc orderedChars) size() int {
	return oc.charList.Len()
}

type ocList struct {
	ocs *list.List
}

func newOcl() *ocList {
	result := new(ocList)
	result.ocs = list.New()
	return result
}

func (ocl *ocList) strRep() string {
	result := "{"
	mark := ocl.ocs.Front()
	for mark != nil {
		oc := mark.Value.(*orderedChars)
		result = result + oc.strRep() + " "
		mark = mark.Next()
	}
	result = result + "}"
	return result
}

func (ocl *ocList) insertOc(oc *orderedChars) {
	mark := ocl.ocs.Front()
	found := false
	for !found && mark != nil {
		if (mark.Value.(*orderedChars)).size() > oc.size() {
			mark = mark.Next()
		} else {
			found = true
		}
	}
	if mark == nil {
		ocl.ocs.PushBack(oc) // empty case as well as end of the list
	} else {
		ocl.ocs.InsertBefore(oc, mark)
	}
}

func (ocl *ocList) add(extra *ocList) {
	mark := extra.ocs.Front()
	for mark != nil {
		ocl.insertOc(mark.Value.(*orderedChars))
		mark = mark.Next()
	}
}

func (ocl *ocList) consumeATuple(t tuple) (*ocList, error) {
	var err error
	extraOcl := newOcl()
	subsumed := false
	notfound := true
	mark := ocl.ocs.Front()
	for mark != nil {
		anOc := mark.Value.(*orderedChars)
		oc, found, err := anOc.consume(t)
		if err != nil {
			break
		}
		if oc != nil {
			extraOcl.ocs.PushBack(oc)
		}
		switch found {
		case 0:
			notfound = notfound && true
		case 3:
			subsumed = subsumed || true
		default:
			notfound = notfound && false
		}
		mark = mark.Next()
	}
	if notfound && !subsumed {
		extraOcl.ocs.PushBack(newOcForTuple(t))
	}
	return extraOcl, err
}

func (ocl *ocList) consume(ts []tuple) {
	if ocl.ocs.Len() == 0 {
		// nothing is added so far, we simply absorb all tuples
		t := ts[0]
		oc := newOcForTuple(t)
		ocl.insertOc(oc)
	} else {
		for _, t := range ts {
			extra, err := ocl.consumeATuple(t)
			if err != nil {
				return
			} else {
				ocl.add(extra)
			}
		}
		// Now that we have consumed a single tuple, we need to see what
		// consolidation of ocs can happen. For example, when we consume
		// tuple (e,r) for ocs [we] and [rtf], we get ocs are [wer] and
		// [ertf] which can be combined into one single oc [wertf].
		// Similarly we have to check if any one oc becomes a subset of
		// another oc and it can be subsumed.
	}
}

func generateTuples(smaller rune, biggers []rune) []tuple {
	var tlist []tuple
	for _, v := range biggers {
		t := new(tuple)
		t.smallerChar = smaller
		t.biggerChar = v
		tlist = append(tlist, *t)
	}
	return tlist
}

func compare(w1 string, w2 string) {
	matching := true
	i := 0
	for matching {
		if i < len(w1) && i < len(w2) && w1[i] == w2[i] {
			i++
		} else {
			matching = false
		}
	}
	if !matching && i < len(w1) && i < len(w2) {
		firstChar := rune(w1[i])
		secondChar := rune(w2[i])
		arr := greaterList[firstChar]
		if arr != nil {
			greaterList[firstChar] = append(arr, secondChar)
		} else {
			arr = make([]rune, 1)
			arr[0] = secondChar
			greaterList[firstChar] = arr
		}
	}
}

var greaterList map[rune][]rune
var ocLists *ocList

func alienOrder(words []string) string {
	greaterList = make(map[rune][]rune)
	w := words[0]
	for _, aw := range words[1:] {
		fmt.Println(fmt.Sprintf("Comparing %v with %v", w, aw))
		compare(w, aw)
		w = aw
	}
	fmt.Println(dumpMap(greaterList))

	ocLists := newOcl()
	for r, rl := range greaterList {
		tuples := generateTuples(r, rl)
		fmt.Println(dumpTupleSlice(tuples))
		ocLists.consume(tuples)
		fmt.Println(ocLists.strRep())
	}

	anOc := ocLists.ocs.Front().Value.(*orderedChars)
	return anOc.strRep()
}

func main() {
	input := []string{"wrt",
		"wrf",
		"er",
		"ett",
		"rftt"}

	fmt.Println(alienOrder(input))

	//input = []string{"z","x"}
	//fmt.Println(alienOrder(input))
}

func dumpMap(m map[rune][]rune) string {
	result := ""
	for r, s := range m {
		result = result + string(r) + ": " + string(s) + "\n"
	}
	return result
}

func dumpTupleSlice(ts []tuple) string {
	result := "{"
	for _, t := range ts {
		result = result + t.string()
	}
	return result + "}"
}
