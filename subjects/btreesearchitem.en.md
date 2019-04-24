# btreeinsertdata

## Instructions

Write a function that searches for an item with a data element equal to elem and return that node

## Expected function

```go
func BTreeSearchItem(root *piscine_test.TreeNode, elem string) *piscine_test.TreeNode {

}

```

## Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
       "fmt"
       piscine "."
)

func main() {
	root := &piscine_test.TreeNode{Data: "4"}
	piscine_test.BTreeInsertData(root, "1")
	piscine_test.BTreeInsertData(root, "7")
	piscine_test.BTreeInsertData(root, "5")
	selected := BTreeSearchItem(root, "7")
	fmt.Print("Item selected -> ")
	if selected != nil {
		fmt.Println(selected.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Parent of selected item -> ")
	if selected.Parent != nil {
		fmt.Println(selected.Parent.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Left child of selected item -> ")
	if selected.Left != nil {
		fmt.Println(selected.Left.Data)
	} else {
		fmt.Println("nil")
	}

	fmt.Print("Right child of selected item -> ")
	if selected.Right != nil {
		fmt.Println(selected.Right.Data)
	} else {
		fmt.Println("nil")
	}
}
```

And its output :

```console
student@ubuntu:~/piscine/btreesearchitem$ go build
student@ubuntu:~/piscine/btreesearchitem$ ./btreesearchitem
Item selected -> 7
Parent of selected item -> 4
Left child of selected item -> 5
Right child of selected item -> nil
student@ubuntu:~/piscine/btreesearchitem$
```
