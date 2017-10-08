AVL Tree Implementation in Go

AVL Tree Package for go lang.

Any data type can be stored in the tree as long as that type implements the following
Inteface:

// AvlTreeValue : Values stored in the AvlTree must implement this interface
type AvlTreeValue interface {
	//Compares the current value with the paramter and returns
	//avl.GT (if the current value is Greater Than the parameter)
	//avl.LT (if the current value is Less than the parameter)
	//avl.EQ (if the current value is Equal to the parameter)
	Compare(AvlTreeValue) uint8
}
