// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

// Another set implementation.
type IntegerSet struct {
	intBoolMap map[int]bool
}

func NewSet() *IntegerSet {
	return &IntegerSet{make(map[int]bool)}
}

func (set *IntegerSet) Add(i int) bool {
	_, found := set.intBoolMap[i]
	set.intBoolMap[i] = true
	return !found //False if it existed already
}

func (set *IntegerSet) Get(i int) bool {
	_, found := set.intBoolMap[i]
	return found //true if it existed already
}

func (set *IntegerSet) Remove(i int) {
	delete(set.intBoolMap, i)
}
