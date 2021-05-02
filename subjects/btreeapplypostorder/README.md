## btreeapplypostorder

### Instructions

Write a function that applies a function using a postorder walk to each element in the tree.

### Expected function

```go
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {

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
	piscine.BTreeApplyPostorder(root, fmt.Println)

}
```

And its output :

```console
$ go run .
1
5
7
4
$
```
