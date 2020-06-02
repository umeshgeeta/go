/*
 * MIT License
 * Author: Umesh Patil, Neosemantix, Inc.
 */
package main

type IntSet struct {
	intBoolMap map[int]bool
}

func NewSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.intBoolMap[i]
	set.intBoolMap[i] = true
	return !found //False if it existed already
}

func (set *IntSet) Get(i int) bool {
	_, found := set.intBoolMap[i]
	return found //true if it existed already
}

func (set *IntSet) Remove(i int) {
	delete(set.intBoolMap, i)
}
