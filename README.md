# AVL Tree Implementation in Go
AVL Tree Package for go lang.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine 
for development and testing purposes.

Pull this AVL Package for use in your GO PATH:

```
go get github.com/toravir/go-avl
```

## How to Use
Any data type can be stored in the tree as long as that type implements the following
Inteface:
```
// AvlTreeValue : Values stored in the AvlTree must implement this interface
type AvlTreeValue interface {
	//Compares the current value with the paramter and returns
	//avl.GT (if the current value is Greater Than the parameter)
	//avl.LT (if the current value is Less than the parameter)
	//avl.EQ (if the current value is Equal to the parameter)
	Compare(AvlTreeValue) uint8
}
```

APIs Supported as of now:

```
Create New AVL Tree:
    func NewAvlTree() *TreeNode

Insert Data into AVL Tree:
    func InsertVal(val AvlTreeValue, node *TreeNode) *TreeNode 

Lookup of a data in the AVL Tree:
    func (node *TreeNode) LookupVal(val AvlTreeValue) bool

Delete a Previously inserted data from AVL Tree:
    func DeleteVal(val AvlTreeValue, node *TreeNode) (*TreeNode, error)

```

## Examples

Please refer to the avl_test.go - which has unit-test code.

