## btreedeletenode

### Instructions

Write a function, `BTreeDeleteNode`, that deletes `node` from the tree given by `root`.

The resulting tree should still follow the binary search tree rules.

### Expected function

```go
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	node := piscine.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Before delete:
1
4
5
7
After delete:
1
5
7
student@ubuntu:~/piscine-go/test$
```
