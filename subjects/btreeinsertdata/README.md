## btreeinsertdata

### Instructions

Write a function that inserts new data in a `binary search tree` following the special properties of a `binary search trees`.

The nodes must be defined as follows :

### Notions

- [binary search trees](https://en.wikipedia.org/wiki/Binary_search_tree)

### Expected function

```go
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

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
	fmt.Println(root.Left.Data)
	fmt.Println(root.Data)
	fmt.Println(root.Right.Left.Data)
	fmt.Println(root.Right.Data)

}
```

And its output :

```console
$ go run .
1
4
5
7
$
```
