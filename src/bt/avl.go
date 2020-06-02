/*
 * MIT License
 * Author: Umesh Patil, Neosemantix, Inc.
 */

// AVL Tree - Balanced Binary Search Tree deploying AVL algorithm to balance.
// https://algorithms.tutorialhorizon.com/avl-tree-insertion/

package bt

// *************************************
// Struct & Constant definitions
// *************************************
type Avl struct {
	Bst
}

// *************************************
// Constructor
// *************************************

func NewAvl() *Avl {
	bint := NewBst()
	avl := Avl{*bint}
	return &avl
}

// *************************************
// Methods - Exposed
// *************************************

func (avl *Avl) Insert(val int) error {
	n := NumVal{val}
	nd := Node{
		Value:      n,
		Parent:     nil,
		LeftChild:  nil,
		RightChild: nil,
		LeftDepth:  0,
		RightDept:  0,
	}
	//fmt.Printf("Created node: %v\n", nd)
	return avl.InsertNode(avl.Root, nd)
}

func (avl *Avl) InsertNode(parent *Node, nd Node) error {
	var err error = nil
	if parent == nil {
		// node itself becomes the root
		nd.LeftDepth = 0
		nd.RightDept = 0
		avl.Root = &nd
		avl.NodeCount = 1
	} else {
		comp, _ := parent.Value.Compare(nd.Value)
		switch comp {
		case 0:
			lc := parent.LeftChild
			if lc != nil {
				err = avl.InsertNode(lc, nd)
			} else {
				err = avl.SetChild(&nd, parent, 1, false)
			}
		case 1:
			lc := parent.LeftChild
			if lc != nil {
				err = avl.InsertNode(lc, nd)
			} else {
				err = avl.SetChild(&nd, parent, 1, false)
			}
		case -1:
			rc := parent.RightChild
			if rc != nil {
				err = avl.InsertNode(rc, nd)
			} else {
				err = avl.SetChild(&nd, parent, 0, false)
			}
		}
		balDiff := parent.LeftDepth - parent.RightDept

		if balDiff > 1 && parent.LeftChild != nil {
			com, _ := nd.Value.Compare(parent.LeftChild.Value)
			if com > 0 {
				nl := avl.leftRotate(parent.LeftChild)
				parent.LeftChild = &nl
			}
			if com != 0 {
				avl.rightRotate(parent)
			}
		}

		if balDiff < -1 && parent.RightChild != nil {
			com, _ := nd.Value.Compare(parent.RightChild.Value)
			if com < 0 {
				nl := avl.rightRotate(parent.RightChild)
				parent.RightChild = &nl
			}
			if com != 0 {
				avl.leftRotate(parent)
			}
		}
	}
	return err
}

// *************************************
// Methods - Not Exposed
// *************************************

func (avl *Avl) rightRotate(y *Node) Node {
	x := y.LeftChild
	t2 := x.RightChild
	parent := y.Parent

	// Rotation
	y.Parent = x
	x.RightChild = y

	if parent != nil {
		wc := parent.whichChild(y)
		x.Parent = parent
		switch wc {
		case 0:
			parent.RightChild = x
		case 1:
			parent.LeftChild = x
		}
	} else {
		// else y was root, now x will be root
		x.Parent = nil
		avl.Root = x
	}

	y.LeftChild = t2
	if t2 != nil {
		t2.Parent = y
		y.LeftDepth = t2.GetMaxDepth() + 1
	} else {
		y.LeftDepth = 0
	}
	x.RightDept = y.GetMaxDepth() + 1

	return *x
}

func (avl *Avl) leftRotate(x *Node) Node {
	y := x.RightChild
	t2 := y.LeftChild
	parent := x.Parent

	// Rotation
	x.Parent = y
	y.LeftChild = x

	if parent != nil {
		wc := parent.whichChild(x)
		y.Parent = parent
		switch wc {
		case 0:
			parent.RightChild = y
		case 1:
			parent.LeftChild = y
		}
	} else {
		// x was root, now y becomes the root
		y.Parent = nil
		avl.Root = y
	}

	x.RightChild = t2
	if t2 != nil {
		t2.Parent = x
		x.RightDept = t2.GetMaxDepth() + 1
	} else {
		x.RightDept = 0
	}
	y.LeftDepth = x.GetMaxDepth() + 1

	return *y
}
