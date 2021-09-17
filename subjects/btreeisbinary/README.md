## btreeisbinary

### Instructions

Write a function, `BTreeIsBinary`, that returns `true` only if the tree given by `root` follows the binary search tree [properties](https://en.wikipedia.org/wiki/Binary_search_tree#Definition).

### Expected function

```go
func BTreeIsBinary(root *TreeNode) bool {

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
	fmt.Println(piscine.BTreeIsBinary(root))
}
```

And its output :

```console
$ go run .
true
$
```
