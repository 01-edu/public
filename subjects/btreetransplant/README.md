## btreetransplant

### Instructions

In order to move subtrees around within the binary search tree, write a function, `BTreeTransplant`, which replaces the subtree started by `node` with the node `rplc` in the tree given by `root`.

### Expected function

```go
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

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
	node := piscine.BTreeSearchItem(root, "1")
	rplc := &piscine.TreeNode{Data: "3"}
	root = piscine.BTreeTransplant(root, node, rplc)
	piscine.BTreeApplyInorder(root, fmt.Println)
}
```

And its output :

```console
$ go run .
3
4
5
7
$
```
