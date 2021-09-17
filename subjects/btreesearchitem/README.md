## btreesearchitem

### Instructions

Write a function that returns the `TreeNode` with a `data` field equal to `elem` if it exists in the tree, otherwise return `nil`.

### Expected function

```go
func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

import (
	"fmt"
	"piscine"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	selected := piscine.BTreeSearchItem(root, "7")
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
$ go run .
Item selected -> 7
Parent of selected item -> 4
Left child of selected item -> 5
Right child of selected item -> nil
$
```
