package avl

import (
	"fmt"
	"testing"
)

type testInt int

func (t testInt) Compare(v AvlTreeValue) uint8 {
	vint := v.(testInt)
	if t > vint {
		return GT
	} else if t < vint {
		return LT
	}
	return EQ
}

//Run tests with a int based Type
func TestAvlTreeInt(t *testing.T) {
	root := NewAvlTree()

	root = InsertVal(testInt(1), root)
	root = InsertVal(testInt(2), root)
	q1 := root.printTreeToString()
	e1 := "1/1,2/0"
	if q1 != e1 {
		t.Fail()
		fmt.Println("1. Expected:", e1, ".instead got:", q1, ":")
	}
	root = InsertVal(testInt(3), root) // Triggers a RR case
	e2 := "1/0,2/1,3/0"
	q1 = root.printTreeToString()
	if q1 != e2 {
		t.Fail()
		fmt.Println("2. Expected:", e2, ".instead got:", q1, ":")
	}

	root = InsertVal(testInt(30), root)
	root = InsertVal(testInt(10), root) // RL case
	e3 := "1/0,2/2,3/0,10/1,30/0"
	q1 = root.printTreeToString()
	if q1 != e3 {
		t.Fail()
		fmt.Println("3. Expected:", e3, ".instead got:", q1, ":")
	}

	root = InsertVal(testInt(-10), root)
	root = InsertVal(testInt(-20), root) // LL case

	e4 := "-20/0,-10/1,1/0,2/2,3/0,10/1,30/0"
	q1 = root.printTreeToString()
	if q1 != e4 {
		t.Fail()
		fmt.Println("4. Expected:", e4, ".instead got:", q1, ":")
	}

	root = InsertVal(testInt(-5), root)
	root = InsertVal(testInt(-3), root) // LR case

	e5 := "-20/0,-10/2,-5/0,-3/1,1/0,2/3,3/0,10/1,30/0"
	q1 = root.printTreeToString()
	if q1 != e5 {
		t.Fail()
		fmt.Println("5. Expected:", e5, ".instead got:", q1, ":")
	}

	//root.printTree()
	//fmt.Println("")

	var rc error
	//Deletion
	root, rc = DeleteVal(testInt(100), root)
	if rc == nil {
		t.Fail()
		fmt.Println("5a. Incorrect status returned when attempting to delete non-existing node")
	}
	root, rc = DeleteVal(testInt(30), root)

	e6 := "-20/0,-10/2,-5/0,-3/1,1/0,2/3,3/0,10/1"
	q1 = root.printTreeToString()
	if q1 != e6 {
		t.Fail()
		fmt.Println("6. Expected:", e6, ".instead got:", q1, ":")
	}

	root, rc = DeleteVal(testInt(3), root) //LR Case

	e7 := "-20/0,-10/1,-5/0,-3/2,1/0,2/1,10/0"
	q1 = root.printTreeToString()
	if q1 != e7 {
		t.Fail()
		fmt.Println("7. Expected:", e7, ".instead got:", q1, ":")
	}

	root, rc = DeleteVal(testInt(-3), root)

	e8 := "-20/0,-10/1,-5/0,1/2,2/1,10/0"
	q1 = root.printTreeToString()
	if q1 != e8 {
		t.Fail()
		fmt.Println("8. Expected:", e8, ".instead got:", q1, ":")
	}

	l1 := root.LookupVal(testInt(123))
	if l1 == true {
		t.Fail()
		fmt.Println("9. Expected fail for lookup of non-inserted value..")
	}
	l1 = root.LookupVal(testInt(10))
	if l1 == false {
		t.Fail()
		fmt.Println("10. Expected fail for lookup of inserted (10) value..")
	}
}

type testString string

func (t testString) Compare(v AvlTreeValue) uint8 {
	vstr := v.(testString)
	if t > vstr {
		return GT
	} else if t < vstr {
		return LT
	}
	return EQ
}

//Run tests with a string based Type
func TestAvlTreeString(t *testing.T) {
	root := NewAvlTree()

	root = InsertVal(testString("bcd"), root)
	root = InsertVal(testString("bcdc"), root)
	q1 := root.printTreeToString()
	e1 := "bcd/1,bcdc/0"
	if q1 != e1 {
		t.Fail()
		fmt.Println("S1. Expected:", e1, ".instead got:", q1, ":")
	}
	root = InsertVal(testString("cbcd"), root) // Triggers a RR case
	e2 := "bcd/0,bcdc/1,cbcd/0"
	q1 = root.printTreeToString()
	if q1 != e2 {
		t.Fail()
		fmt.Println("S2. Expected:", e2, ".instead got:", q1, ":")
	}

	root = InsertVal(testString("kjl"), root)
	root = InsertVal(testString("fgh"), root) // RL case
	e3 := "bcd/0,bcdc/2,cbcd/0,fgh/1,kjl/0"
	q1 = root.printTreeToString()
	if q1 != e3 {
		t.Fail()
		fmt.Println("S3. Expected:", e3, ".instead got:", q1, ":")
	}

	root = InsertVal(testString("aman"), root)
	root = InsertVal(testString("abc"), root) // LL case

	e4 := "abc/0,aman/1,bcd/0,bcdc/2,cbcd/0,fgh/1,kjl/0"
	q1 = root.printTreeToString()
	if q1 != e4 {
		t.Fail()
		fmt.Println("S4. Expected:", e4, ".instead got:", q1, ":")
	}

	root = InsertVal(testString("axyz"), root)
	root = InsertVal(testString("azzz"), root) // LR case

	e5 := "abc/0,aman/2,axyz/0,azzz/1,bcd/0,bcdc/3,cbcd/0,fgh/1,kjl/0"
	q1 = root.printTreeToString()
	if q1 != e5 {
		t.Fail()
		fmt.Println("S5. Expected:", e5, ".instead got:", q1, ":")
	}

	var rc error
	//Deletion
	root, rc = DeleteVal(testString("foobar"), root)
	if rc == nil {
		t.Fail()
		fmt.Println("5a. Incorrect status returned when attempting to delete non-existing node")
	}
	root, rc = DeleteVal(testString("kjl"), root)

	e6 := "abc/0,aman/2,axyz/0,azzz/1,bcd/0,bcdc/3,cbcd/0,fgh/1"
	q1 = root.printTreeToString()
	if q1 != e6 {
		t.Fail()
		fmt.Println("S6. Expected:", e6, ".instead got:", q1, ":")
	}

	root, rc = DeleteVal(testString("cbcd"), root) //LR Case

	e7 := "abc/0,aman/1,axyz/0,azzz/2,bcd/0,bcdc/1,fgh/0"
	q1 = root.printTreeToString()
	if q1 != e7 {
		t.Fail()
		fmt.Println("S7. Expected:", e7, ".instead got:", q1, ":")
	}

	root, rc = DeleteVal(testString("azzz"), root)

	e8 := "abc/0,aman/1,axyz/0,bcd/2,bcdc/1,fgh/0"
	q1 = root.printTreeToString()
	if q1 != e8 {
		t.Fail()
		fmt.Println("S8. Expected:", e8, ".instead got:", q1, ":")
	}
}
