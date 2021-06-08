## btreeapplyinorder

### Instructions

Write a function that applies a given function `f`, **in order**, to each element in the tree.

### Notions

- [Tree Traversal](https://en.wikipedia.org/wiki/Tree_traversal)

### Expected function

```go
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {

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
	piscine.BTreeApplyInorder(root, fmt.Println)

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
