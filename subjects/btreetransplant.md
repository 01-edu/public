## btreetransplant
### Instructions

In order to move subtrees around within the binary search tree, write a function, BTreeTransplant, which replaces the subtree started by node with the node called 'rplc' in the tree given by root.

This function must have the following signature.

### Expected function

```go
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	
}

```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
       "fmt"
       student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "1")
	replacement := &student.TreeNode{Data: "3"}
	root = student.BTreeTransplant(root, node, replacement)
	student.BTreeApplyInorder(root, fmt.Println)
}
```

And its output :

```console
student@ubuntu:~/student/btreetransplant$ go build
student@ubuntu:~/student/btreetrandsplant$ ./btreetransplant
3
4
5
7
student@ubuntu:~/student/btreetrandsplant$ 
```
