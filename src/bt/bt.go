// author: Umesh Patil
// March 2020

// The program builds a binary tree based on insertions done. The level
// where root resides is designated as '1', next at '2' and so forth.
// If height of the binary tree is 'h'; there will be h levels in the tree.
//
//
// It prints the binary tree on console as a pyramid. It computes
// spaces so that the pyramid is depicted as closely as possible.
package bt

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

// *************************************
// Struct & Constant definitions
// *************************************
type Comparable interface {
	// if 'this' equals c2:	returns 0
	// if 'this' is greater than c2: returns 1
	// if 'this' is smaller than c2: returns -1
	// else error
	Compare(c2 Comparable) (int, error)

	// returns the length of the value for printing purposes
	Len() int

	// Returns the string representation of this value
	Str() string
}

type NumVal struct {
	Val int
}

type Bt struct {
	Root *Node
	NodeCount int
}

type Node struct {
	Value Comparable
	Parent *Node
	LeftChild *Node
	RightChild *Node
	LeftDepth int
	RightDept int
}

const TileSize = 4
const PadChar = " "
const TileChar = " "

// *************************************
// Constructor
// *************************************
func NewBt() *Bt {
	bt := new(Bt)
	return bt
}

func NewBtWithRoot(rv Comparable) *Bt {
	bt := new(Bt)
	root := Node {rv, nil, nil, nil, 0, 0}
	bt.SetChild(&root, nil, 0)
	return bt
}

func NewBtWithNumValRoot(val int) *Bt {
	nv := NumVal{val}
	return NewBtWithRoot(nv)
}

// *************************************
// Methods - Exposed
// *************************************

// Interface implementation, we introduce a type NumVal as Comparable.
// Below is the implementation of how that comparision should be done.
func (nv NumVal) Compare(c2 Comparable) (int, error) {
	result := 0
	var err error
	f, ok := c2.(NumVal)
	if ok {
		if nv.Val > f.Val {
			result = 1
		} else if nv.Val < f.Val {
			result = -1
		}
		// else existing 0 value is good
	} else {
		err = errors.New("Argument Comparable is not of type NumVal")
	}
	return result, err
}

func (nv NumVal) Len() int {
	return len(nv.Str())
}

func (nv NumVal) Str() string {
	return strconv.Itoa(nv.Val)
}

// byWhat: It can be +1 or -1 or some other when the entire subtree is appended or chopped off
// whichChild: 1 for Left Child and 0: for Right Child
func (node* Node) ChangeDepth(byWhat int, whichChild int) {
	if whichChild == 1 {
		node.LeftDepth += byWhat
		if node.Parent != nil {
			// recursive call
			node.Parent.ChangeDepth(byWhat, whichChild)
		}
	} else if whichChild == 0 {
		node.RightDept += byWhat
		if node.Parent != nil {
			// recursive call
			node.Parent.ChangeDepth(byWhat, whichChild)
		}
	} else {
		log.Fatal(fmt.Sprintf("Wrong depth argument %d for increaseDepth call", whichChild))
	}
}

// Throw an error if the designated place is already occupied.
// Argument lr indicates whether to set this node as left child (lr == 0)or
// right child (lr == 1) of the parent. When parent is nil, we treat
// invocation of the method as to set the root in which case argument lr is ignored.
func (bt* Bt) SetChild(node *Node, parent *Node, lr int) error {
	var err error
	err = nil
	// if parent is nil and the tree does not have any root, we can place the node at the root
	if parent == nil {
		if bt.Root == nil {
			bt.Root = node
			bt.NodeCount = 1
			node.LeftDepth = 0
			node.RightDept = 0
		} else {
			err = errors.New("Root already exists, cannot set.")
		}
	} else {
		switch lr {
		case 0:
			if parent.RightChild == nil {
				parent.RightChild = node
				node.Parent = parent
				bt.NodeCount++
				node.ChangeDepth(1, 0)
			} else {
				err = errors.New("Right child already exists, cannot set.")
			}
		case 1:
			if parent.LeftChild == nil {
				parent.LeftChild = node
				node.Parent = parent
				node.ChangeDepth(1, 1)
				bt.NodeCount++
			} else {
				err = errors.New("Left child already exists, cannot set.")
			}
		}
	}
	return err
}

func (bt* Bt) SetNode(index int, node *Node) error {
	var err error
	path := GetPath(index)
	levels := path.Len()
	if levels > 0 {
		e := path.Front()
		path.Remove(e)	// remove the root from the list
		// we expect the first node to be root and value at 0
		child := e.Value.(int)
		if child == 0 {
			level := 1
			parent := bt.Root
			for e = path.Front(); e != nil && err == nil; e = e.Next() {
				child = e.Value.(int)
				level++
				if level < levels {
					if child == 0 {
						if parent != nil {
							// even - it is Right Child
							parent = parent.RightChild
						} else {
							err = errors.New(fmt.Sprintf("No parent node at level %d to attach right child", level))
						}
					} else if child == 1 {
						if parent != nil {
							// odd - it is Left Child
							parent = parent.LeftChild
						} else {
							err = errors.New(fmt.Sprintf("No parent node at level %d to attach left child", level))
						}
					} else {
						err = errors.New(fmt.Sprintf("Unexpected path value %w", e.Value))
					}
				}
				// else the last level is where the node is to be placed actually
			}
			// we set it as the Child
			bt.SetChild(node, parent, child)
		} else {
			err = errors.New("Invalid path, did not find expected root node at the start")
		}
	} else {
		err = errors.New(fmt.Sprintf("No vaiable path available for index %d", index))
	}
	return err
}

func (bt* Bt) Insert(index int, val int) error {
	n := NumVal{val}
	nd := Node{
		Value:      n,
		Parent:     nil,
		LeftChild:  nil,
		RightChild: nil,
		LeftDepth:  0,
		RightDept:  0,
	}
	return bt.SetNode(index, &nd)
}

func (bt Bt) Height() int {
	h := 0
	if bt.Root != nil {
		depth := int(math.Max(float64(bt.Root.LeftDepth), float64(bt.Root.RightDept)))
		h = depth + 1	// we need to add the root
	}
	return h
}

// Returns the tree as a list upon traversing as Breadth First Search (BFS).
// Where a node is absent, 'nil' is added as the element.
func (bt Bt) Bfs() (list.List, int) {
	l := list.New()
	maxlen := 0
	if bt.Root != nil {
		max := bt.MaxNodesPossible()
		nodesAdded := 0
		e := l.PushBack(bt.Root)
		maxlen = bt.Root.Value.Len()
		nodesAdded++
		var pn *Node
		var chl *Node
		var chr *Node
		for nodesAdded < max {
			if e.Value != nil {
				pn = e.Value.(*Node)
				if pn != nil {
					chl = pn.LeftChild
					chr = pn.RightChild
					ln := pn.Value.Len()
					if maxlen < ln {
						maxlen = ln
					}
				} else {
					chl = &Node{nil, nil,nil,nil,0, 0}
					chr = &Node{nil, nil,nil,nil,0, 0}
					l.PushBack(chl)
					l.PushBack(chr)
				}
				l.PushBack(chl)
				l.PushBack(chr)
				nodesAdded += 2

			} else {
				// wrap the nil
				n := Node{nil, nil,nil,nil,0, 0}
				//l.InsertAfter(n, e)
				l.PushBack(&n)
			}
			e = e.Next()
		}
	}
	// essentially it is all empty tree, so we return empty list
	return *l, maxlen
}

// Returns the pyramid string which can be printed.
// Gaps are called as tiles. Numbers are padded to maximum value so as alignment works.
func (bt Bt) Pyramid() string {
	result := ""
	th := bt.Height()
	tpr := tilesPerRow(th)
	bfs, maxl := bt.Bfs()
	where := bfs.Front()
	for level := 1; level <= th; level++ {
		td := tileDist(tpr, level, th)
		str, wh :=  bt.fillRow(level, td, bfs, where, maxl)
		result  = result + str + "\n"
		where = wh
	}
	return result
}

// *************************************
// Methods - Not Exposed
// *************************************

// Fills the row for the given level
func (bt Bt) fillRow(level int, td []int, bfs list.List, where *list.Element, maxlen int) (string, *list.Element) {
	result := ""
	nodesInLevel := numPerLevel(level)
	charsToAdjust := (bt.maxNodesInARow() - nodesInLevel) * maxlen
	charsPerNode := charsToAdjust / nodesInLevel
	tdl := len(td)
	for k := 0; k < tdl-1; k++ {
		result += tileStr(td[k])
		valstr := "nil"
		if where != nil {
			nd := where.Value.(*Node)
			if nd != nil {
				ndv := nd.Value
				if ndv != nil {
					valstr = ndv.Str()
				}
			}
		}
		result += bt.padNode(valstr, charsPerNode, maxlen)
		where = where.Next()
	}
	result += tileStr(td[tdl-1])
	return result, where
}

// Pad the value at a node so as it aligns to the value with maximum digits.
func (bt Bt) padNode(val string, additional int, maxlen int) string {
	return pad(val, maxlen, additional)
}

// Note that it is NOT how many'nodes' are present in the tree;
// but how many can be 'effectively' packed in this binary tree.
func (bt Bt) MaxNodesPossible() int {
	return int(math.Pow(2, float64(bt.Height()))) - 1
}

// Indicates how many node - values - can be maximum in a row,
// that row is the last row or the last level.
func (bt Bt) maxNodesInARow() int {
	return numPerLevel(bt.Height())
}

//func (bt Bt) setChild(parent int, childIndex int, childVal int) bool {
//	result := false
//	if childIndex < bt.nodeCount() {
//		bt.Insert(childIndex,childVal)
//		result = true
//	}
//	return result
//}

// *************************************
// Functions - Exposed
// *************************************

func GetPath(index int) *list.List {
	l := list.New()
	i := index
	for i >= 0 {
		r := i % 2
		l.PushFront(r)
		switch r {
		case 0:
			i = i/2 - 1
		case 1:
			i = (i-1) / 2
		}
	}
	return l
}


// *************************************
// Functions - Not Exposed
// *************************************

// returns number of nodes for a given level
func numPerLevel(level int) int {
	return int(math.Pow(2, float64(level-1)))
}

// start and end indices of nodes in a given level
// end index is excluded, upto that value
func indices(l int) (int, int){
	start := int(math.Pow(2, float64(l-1))) - 1
	end := start + int(math.Pow(2, float64(l-1)))
	return start, end
}

// We pad the value to max. digit length, plus additionally
// more padding because the number of values at that level will
// be less than the maximum. For a binary tree with height 4,
// the last row / level will have 8 numbers which second last
// will have only 4 numbers. So the padding for those 4 additional
// numbers on the last row needs to be distributed among 4 numbers
// when we deal with padding of numbers on level 3.
func pad(ns string, ps int, additional int) string {
	d := len(ns)
	result := ""
	totalChars := additional + ps - d
	half := totalChars / 2
	for j := 0; j < half; j++ {
		result += PadChar
	}
	result += ns
	for j := half; j < totalChars; j++ {
		result += PadChar
	}
	return result
}

func tileStr(tiles int) string {
	result := ""
	if tiles > 0 {
		for k := 0; k < tiles; k++ {
			result += singleTile()
		}
	}
	return result
}

func singleTile() string {
	result := ""
	for i := 0; i < TileSize; i++ {
		result += TileChar
	}
	return result
}

func tilesPerRow(height int) int {
	return (int(math.Pow(2, float64(height)))) - 2
}

// Method returns tile distribution for a row of the specified level.
// Key thing to understand is 'second row' in a binary tree of depth /
// height 4 will have different count of tiles compared to the same
// 'second row' in a binary tree of depth 8. That is because the last
// row will have more numbers for a taller binary tree and we want to
// start with 2 tile separation between 2 numbers at the bottom row.
func tileDist(tileNum int, level int, btHeight int) []int {
	tileSeg := int(math.Pow(2, float64(level-1))) + 1
	result := make([]int, tileSeg)
	boundary := int(math.Pow(2, float64(btHeight-level))) - 1
	result[0] = boundary
	result[tileSeg-1] = boundary
	tilesLeft := tileNum - (boundary * 2)
	if tilesLeft > 0 {
		perSeg := tilesLeft / (tileSeg - 2)
		for k := 1; k < tileSeg-1; k++ {
			result[k] = perSeg
		}
		leftOver := tilesLeft % (tileSeg - 2)
		if leftOver > 0 {
			middle := tileSeg / 2
			// for now add all to the middle segment
			// TODO - distribute around
			result[middle] += leftOver
		}
	}
	return result
}