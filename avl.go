/*
Copyright 2017 Ravi Raju <toravir@yahoo.com>

Redistribution and use in source and binary forms, with or without modification, are
permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list
of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this
list of conditions and the following disclaimer in the documentation and/or other
materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT
SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED
TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN
ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
SUCH DAMAGE.

*/

package avl

import (
	"errors"
	"fmt"
)

//These are the return Values for the Comparator function
const (
	Invalid = iota
	EQ
	LT
	GT
)

// AvlTreeValue : Values stored in the AvlTree must implement this interface
type AvlTreeValue interface {
	//Compares the current value with the paramter and returns
	//GT (if the current value is Greater Than the parameter)
	//LT (if the current value is Less than the parameter)
	//EQ (if the current value is Equal to the parameter)
	Compare(AvlTreeValue) uint8
}

//TreeNode to store value
type TreeNode struct {
	key         AvlTreeValue
	refcount    uint //for handling duplicates
	left, right *TreeNode
	height      int
}

//Verbose variable has to be set to TRUE to enable Logging by
//this packages (to STDOUT)
var Verbose = false

func logMessage(a ...interface{}) {
	if Verbose {
		fmt.Println(a...)
	}
}

//NewAvlTree method is invoked to create a blank AvlTree
func NewAvlTree() *TreeNode {
	var blankTree *TreeNode
	blankTree = nil
	return blankTree
}

func updateHeight(node *TreeNode) {
	if node == nil {
		return
	}
	lh, rh := -1, -1
	if node.left != nil {
		lh = node.left.height
	}
	if node.right != nil {
		rh = node.right.height
	}
	if lh > rh {
		node.height = lh + 1
	} else {
		node.height = rh + 1
	}
	return
}

/*
Input tree - pointer to Node '3'

				3
		   +---------+
		   2         A
	   +-------+
	   1       B
   +-------+
   C       D

Output tree - pointer to Node '2'

				  2
			+-----------+
            1           3
		 +------+    +------+
		 C      D    B      A


*/
func rotateRight(node *TreeNode) *TreeNode {
	logMessage("Rotating Right:", node.key)
	tmp := node.left      // tmp = (3).left => 2
	node.left = tmp.right // (3).left = (2).right => B
	tmp.right = node      // (2).right = (3)

	// Adjust the heights correctly
	updateHeight(node)
	updateHeight(tmp)
	return tmp
}

/*
Input tree - pointer to Node '1'

				1
		   +---------+
		   A         2
				 +-------+
				 B       3
                     +--------+
                     C        D

Output tree - pointer to Node '2'

				  2
			+-----------+
            1           3
		 +------+    +------+
		 A      B    C      D


*/
func rotateLeft(node *TreeNode) *TreeNode {
	logMessage("Rotating Left:", node.key)
	tmp := node.right     // tmp = (1).right => 2
	node.right = tmp.left // (1).right = (2).left => B
	tmp.left = node       // (2).left = (1)

	// Adjust the heights correctly
	updateHeight(node)
	updateHeight(tmp)
	return tmp
}

func checkAndBalance(node *TreeNode) *TreeNode {
	updateHeight(node)
	lHeight, rHeight := 0, 0
	if node.left != nil {
		lHeight = node.left.height + 1
	}
	if node.right != nil {
		rHeight = node.right.height + 1
	}
	heightDiff := lHeight - rHeight
	if heightDiff > 1 {
		//this node is left heavy
		chLeftHeight, chRightHeight := -1, -1
		if node.left.left != nil {
			chLeftHeight = node.left.left.height
		}
		if node.left.right != nil {
			chRightHeight = node.left.right.height
		}
		if chLeftHeight > chRightHeight {
			//LL case
			logMessage("LL Case")
			node = rotateRight(node)
		} else {
			//LR case
			logMessage("LR Case")
			node.left = rotateLeft(node.left)
			node = rotateRight(node)
		}
	} else if heightDiff < -1 {
		//This node is right heavy
		chLeftHeight, chRightHeight := -1, -1
		if node.right.left != nil {
			chLeftHeight = node.right.left.height
		}
		if node.right.right != nil {
			chRightHeight = node.right.right.height
		}
		if chLeftHeight < chRightHeight {
			//RR case
			logMessage("RR Case")
			node = rotateLeft(node)
		} else {
			//RL case
			logMessage("RL Case")
			node.right = rotateRight(node.right)
			node = rotateLeft(node)
		}
	}
	return node
}

// InsertVal method is invoked to insert a new value into existing Avl Tree
func InsertVal(val AvlTreeValue, node *TreeNode) *TreeNode {
	if node == nil {
		//We have a blank Tree - so this is the
		//first node
		return &TreeNode{key: val, left: nil, right: nil, height: 0}
	}
	op := val.Compare(node.key)
	if op == EQ {
		//Equal key - we'll incr refcount
		node.refcount++
		return node // we are done
	} else if op == LT {
		//go left
		node.left = InsertVal(val, node.left)
	} else if op == GT {
		//go right
		node.right = InsertVal(val, node.right)
	} else {
		logMessage("Invalid OP returned from Compare:", op)
	}
	return checkAndBalance(node)
}

func findMinkey(node *TreeNode) (AvlTreeValue, uint) {
	for node.left != nil {
		node = node.left
	}
	return node.key, node.refcount
}

func findMaxkey(node *TreeNode) (AvlTreeValue, uint) {
	for node.right != nil {
		node = node.right
	}
	return node.key, node.refcount
}

// DeleteVal method is invoked to Delete a value from an existing Avl Tree
// Returns the new Root and a bool to indicate if
func DeleteVal(val AvlTreeValue, node *TreeNode) (*TreeNode, error) {
	if node == nil {
		logMessage("Element ", val, " NOT Found!")
		return node, errors.New("Element not Found")
	}
	var rc error
	op := val.Compare(node.key)
	if op == GT {
		//go right
		node.right, rc = DeleteVal(val, node.right)
	} else if op == LT {
		//go left
		node.left, rc = DeleteVal(val, node.left)
	} else if op == EQ {
		if node.refcount > 1 {
			node.refcount--
			return node, nil // we are done
		}
		//We have to delete this current Node
		if node.left == nil {
			//leaf node (no children) case AND case with only right tree
			return node.right, nil
			//leave the current Node to be garbage collected..
		} else if node.right == nil {
			return node.left, nil
			//leave the current Node to be garbage collected..
		} else {
			//We have both the children - gotto do either of this:
			//bring inorder predecessor of this node to this place or
			//bring inorder successor   of this node to this place
			//Lets choose the smaller of our children and operate
			var promotedVal AvlTreeValue
			promotedRefs := uint(0)
			if node.left.height < node.right.height {
				promotedVal, promotedRefs = findMaxkey(node.left)
				for i := uint(0); i < promotedRefs+1; i++ {
					node.left, rc = DeleteVal(promotedVal, node.left)
				}
			} else {
				promotedVal, promotedRefs = findMinkey(node.right)
				for i := uint(0); i < promotedRefs+1; i++ {
					node.right, rc = DeleteVal(promotedVal, node.right)
				}
			}
			node.key = promotedVal
			node.refcount = promotedRefs
		}
	} else {
		logMessage("Invalid OP returned from Compare:", op)
	}
	return checkAndBalance(node), rc
}

func (node *TreeNode) printTree() {
	if node == nil {
		fmt.Println("{}")
		return
	}
	//lets do inorder traversal
	if node.left != nil {
		fmt.Print("{")
		node.left.printTree()
		fmt.Print("}")
	} else {
		fmt.Print("{}")
	}
	fmt.Print(node.key, "/", node.height)
	if node.right != nil {
		fmt.Print("{")
		node.right.printTree()
		fmt.Print("}")
	} else {
		fmt.Print("{}")
	}
}

func (node *TreeNode) printTreeToString() string {
	if node == nil {
		return ""
	}
	str := ""
	//lets do inorder traversal
	if node.left != nil {
		str += node.left.printTreeToString()
		str += ","
	} else {
		str += ""
	}
	str += fmt.Sprint(node.key, "/", node.height)
	if node.right != nil {
		str += ","
		str += node.right.printTreeToString()
	} else {
		str += ""
	}
	return str
}

//LookupVal method is invoked to search for a particular value in an Avl Tree
//It returns true if found, false when the val is NOT found in the tree
func (node *TreeNode) LookupVal(val AvlTreeValue) bool {
	if node == nil {
		//Tree is blank Tree..
		return false
	}
	op := val.Compare(node.key)
	if op == GT {
		return node.right.LookupVal(val)
	} else if op == LT {
		return node.left.LookupVal(val)
	}
	return true
}
