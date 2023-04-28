/*
 * LeetCode problem no. 146: LRU Cache
 * https://leetcode.com/problems/lru-cache/description/
 *
 * Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
 *
 * Implement the LRUCache class:
 * LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
 * int get(int key) Return the value of the key if the key exists, otherwise return -1.
 * void put(int key, int value) Update the value of the key if the key exists.
 * Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity
 * from this operation, evict the least recently used key.
 *
 * The functions get and put must each run in O(1) average time complexity.
 *
 * Implementation - basically we use 2 pointers to previous & next key. We need 2 pointers
 * because when we move one of the middle keys to the top - head - we need reference to a
 * key before the one moved to point to the key next in line.
 *
 * MIT License
 * Copyright (c) 2023. Neosemantix, Inc.
 * Author: Umesh Patil
 */

package main

import (
	"sync"
)

type LRUCache struct {
	tailKey  int
	headKey  int
	capacity int
	kvMap    map[int]valueNode
	mux      sync.Mutex
}

type valueNode struct {
	val int
	pkp *int // goes towards Head
	nkp *int // goes towards Tail
}

func Constructor(capacity int) LRUCache {
	lruc := LRUCache{
		tailKey:  0,
		headKey:  0,
		capacity: capacity,
		kvMap:    make(map[int]valueNode),
	}
	return lruc
}

func (this *LRUCache) Get(key int) int {
	result := -1
	existingValueNode, exists := this.kvMap[key]
	if exists {
		this.mux.Lock()
		this.updateToHead(key, existingValueNode, existingValueNode.val)
		result = existingValueNode.val
		this.mux.Unlock()
	}
	return result
}

func (this *LRUCache) Put(key int, value int) {
	this.mux.Lock()
	// first we need to see if a key is present or not
	oldValueNode, exists := this.kvMap[key]
	if !exists {
		if len(this.kvMap) == this.capacity {
			// we need to chop the tail to make a room
			// for the new incoming key
			this.chopTail()
		}
		this.insertHead(key, value)
	} else {
		this.updateToHead(key, oldValueNode, value)
	}
	this.mux.Unlock()
}

func (this *LRUCache) chopTail() {
	tailNode := this.kvMap[this.tailKey]
	tk := this.tailKey
	if tailNode.pkp != nil {
		previousNodeKey := *tailNode.pkp
		previousNode := this.kvMap[previousNodeKey]
		previousNode.nkp = nil // since this is a new tail
		this.kvMap[previousNodeKey] = previousNode
		this.tailKey = previousNodeKey
	}
	delete(this.kvMap, tk)
}

func (this *LRUCache) updateToHead(key int, incomingHeadNode valueNode, newValue int) {
	// if the incoming is already the head, we don't need to change any pointers
	// we simple update the value
	if this.headKey == key {
		crtHeadNode := this.kvMap[key]
		crtHeadNode.val = newValue
		this.kvMap[key] = crtHeadNode
		return
	}

	// we have to check if the oldValueNode was tailKey in which case
	// it will not have next key
	if this.tailKey != key {
		// meaning there are other nodes subsequent to this node
		nextKey := *incomingHeadNode.nkp
		nextKeyNode := this.kvMap[nextKey]
		previousKey := *incomingHeadNode.pkp
		previousKeyNode := this.kvMap[previousKey]

		nextKeyNode.pkp = &previousKey
		previousKeyNode.nkp = &nextKey
		// update the map
		this.kvMap[nextKey] = nextKeyNode
		this.kvMap[previousKey] = previousKeyNode
	} else {
		// else incomingHeadNode was the tailKey, we can simply cut
		if incomingHeadNode.pkp != nil {
			previousKey := *incomingHeadNode.pkp
			previousKeyNode := this.kvMap[previousKey]
			previousKeyNode.nkp = nil
			this.kvMap[previousKey] = previousKeyNode
			// add the new tail key will this previous key
			this.tailKey = previousKey
		}
	}

	incomingHeadNode.val = newValue
	previousHeadKey := this.headKey
	previousHeadNode := this.kvMap[previousHeadKey]
	previousHeadNode.pkp = &key
	// nkp of previousHeadNode remains as is
	this.kvMap[previousHeadKey] = previousHeadNode

	incomingHeadNode.nkp = &previousHeadKey
	incomingHeadNode.pkp = nil
	this.kvMap[key] = incomingHeadNode

	this.headKey = key
}

func (this *LRUCache) insertHead(key int, value int) {
	newHeadNode := valueNode{
		val: value,
		pkp: nil,
		nkp: nil,
	}

	if len(this.kvMap) > 0 {
		// we know for sure that map has more than one entry
		// so there must be current head node
		currentHeadNode := this.kvMap[this.headKey]
		// make the current head node's previous key point to the
		// new incoming key which will be new head key
		currentHeadNode.pkp = &key
		// add back
		this.kvMap[this.headKey] = currentHeadNode
		// now make make next key of the new head key node point
		// to the previous head node key
		previousHeadKey := this.headKey
		newHeadNode.nkp = &previousHeadKey
	}
	// now update the head key
	this.kvMap[key] = newHeadNode
	this.headKey = key

	// if this happens to be the first key,
	// we need to also mark it as the tail key
	if len(this.kvMap) == 1 {
		this.tailKey = key
	}
}

func main() {

	//lRUCache := Constructor(1)
	//lRUCache.Put(2, 1)
	//fmt.Println(lRUCache.Get(2))

	//lRUCache := Constructor(2)
	//lRUCache.Put(1, 0) // cache is {1=1}
	//lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	//lRUCache.Get(1)    // return 1
	//lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	//lRUCache.Get(2)    // returns -1 (not found)
	//lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	//lRUCache.Get(1)    // return -1 (not found)
	//lRUCache.Get(3)    // return 3
	//lRUCache.Get(4)    // return 4

	//lRUCache.Get(2)    // return -1
	//lRUCache.Put(2, 6) // cache is {2=6}
	//lRUCache.Get(1)    // return -1
	//lRUCache.Put(1, 5) // cache is {1=5, 2=6}
	//lRUCache.Put(1, 2) // cache is {1=2, 2=6}
	//lRUCache.Get(1)    // return 2
	//lRUCache.Get(2)    // return 6

	lRUCache := Constructor(3)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	lRUCache.Put(3, 3)
	lRUCache.Put(4, 4)
	lRUCache.Get(4)
	lRUCache.Get(3)
	lRUCache.Get(2)
	lRUCache.Get(1)
	lRUCache.Put(5, 5)
	lRUCache.Get(1)
	lRUCache.Get(2)
	lRUCache.Get(3)
	lRUCache.Get(4)
	lRUCache.Get(5)
}
