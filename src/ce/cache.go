// MIT License
// Author: Umesh Patil, Neosemantix, Inc.

package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
)

const cacheSize = 5

type cache struct {
	keyNodeMap map[string]*list.Element
	nodeList   *list.List
}

type node struct {
	key   string
	value int
}

// Core Get function
func (ch *cache) Get(arg string) int {
	var result int
	np := ch.keyNodeMap[arg]
	if np != nil {
		// we found the node
		result = np.Value.(node).value
		// we need to bring this node to the top
		ch.nodeList.Remove(np)
		front := ch.nodeList.Front()
		ch.nodeList.InsertBefore(np.Value, front)
	} else {
		// we did not find, we need to read from persistence
		newVal := read(arg)
		n := new(node)
		n.value = newVal
		n.key = arg
		var newNode *list.Element
		if len(ch.keyNodeMap) == cacheSize {
			// we know the back node is last used, so we remove that
			last := ch.nodeList.Back()
			ch.nodeList.Remove(last)
			// as well as the key in the map
			lastKey := last.Value.(*node).key
			delete(ch.keyNodeMap, lastKey)
			// we add the new node at the front
			front := ch.nodeList.Front()
			newNode = ch.nodeList.InsertBefore(n, front)
		} else {
			// there is enough place in the cache, we add it
			front := ch.nodeList.Front()

			if front == nil {
				newNode = ch.nodeList.PushFront(n)
			} else {
				// we add it at the front since it is the last used element
				newNode = ch.nodeList.InsertBefore(n, front)
			}
		}
		// add back to the map
		ch.keyNodeMap[arg] = newNode
		result = newVal
	}
	return result
}

// Getting cache value from persistence store....
func read(arg string) int {
	// we do not consume the key, but in real life the key will be consumed to get the value
	// from DB or a filesystem etc.
	// We simply return a random number between 0  and 100 (excluded 100).
	return rand.Intn(100)
}

// For debug purposes
func (ch *cache) dumpCache() string {
	result := "{"
	mark := ch.nodeList.Front()
	for mark != nil {
		pair := "(" + mark.Value.(*node).key + ", " + strconv.Itoa(mark.Value.(*node).value) + ")"
		result = result + pair + ", "
		mark = mark.Next()
	}
	return result + "}"
}

func main() {

	myCache := new(cache)
	myCache.keyNodeMap = make(map[string]*list.Element)
	myCache.nodeList = list.New()

	fmt.Printf("1. %d\n", myCache.Get("key1"))
	fmt.Printf("2. %d\n", myCache.Get("key2"))
	fmt.Printf("3. %d\n", myCache.Get("key3"))
	fmt.Printf("4. %d\n", myCache.Get("key4"))
	fmt.Printf("5. %d\n", myCache.Get("key5"))

	fmt.Println(myCache.dumpCache())

	fmt.Printf("6. %d\n", myCache.Get("key6"))

	fmt.Println(myCache.dumpCache())

	// Expected output is as follows:
	//1. 81
	//2. 87
	//3. 47
	//4. 59
	//5. 81
	//{(key5, 81), (key4, 59), (key3, 47), (key2, 87), (key1, 81), }
	//6. 18
	//{(key6, 18), (key5, 81), (key4, 59), (key3, 47), (key2, 87), }

}
