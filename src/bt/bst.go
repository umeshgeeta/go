// author: Umesh Patil
// March 2020

// Basic implementation of Binary Search Tree. It wraps an instance of a
// binary tree (Bt struct) and the core method is insertion where based on the
// value, the incoming node is added to the Left subtree or Right subtree.

package bt

// *************************************
// Struct & Constant definitions
// *************************************
type Bst struct {
	Bt
}
// *************************************
// Constructor
// *************************************

func NewBst() *Bst {
	bint := NewBt()
	bst1 := Bst { *bint }
	return &bst1
}

// *************************************
// Methods - Exposed
// *************************************

func (bst* Bst)Insert(val int) error {
	n := NumVal{val}
	nd := Node{
		Value:      n,
		Parent:     nil,
		LeftChild:  nil,
		RightChild: nil,
		LeftDepth:  0,
		RightDept:  0,
	}
	var err error
	err = nil
	if bst.Root == nil {
		nd.LeftDepth = 0
		nd.RightDept = 0
		bst.Root = &nd
		bst.NodeCount = 1
	} else {
		parent, whichChild, err := bst.findParent(bst.Root, n)
		if err != nil {
			return err
		}
		err = bst.SetChild(&nd, parent, whichChild)
	}
	return err
}

// *************************************
// Methods - Not Exposed
// *************************************

func (bst* Bst) findParent(root *Node, incoming Comparable) (*Node, int, error) {
	node := root
	var comparison, err = node.Value.Compare(incoming)
	if  err == nil {
		switch comparison {
		case 0, 1:
			if node.LeftChild != nil {
				return bst.findParent(node.LeftChild, incoming)
			} else {
				return node, 1, nil
			}
		case -1:
			if node.RightChild != nil {
				return bst.findParent(node.RightChild, incoming)
			} else {
				return node, 0, nil
			}
		}
	}
	return &Node{}, 0, err
}