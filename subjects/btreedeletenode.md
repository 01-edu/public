# btreedeletenode
## Instructions

Write a function, BTreeDeleteNode, that deletes 'node' from the tree given by root.

The resulting tree should still follow the binary search tree rules.

This function must have the following signature.

## Expected function

```go
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	
}

```

## Usage

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
	node := student.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)
}
```

And its output :

```console
student@ubuntu:~/student/btreedeletenode$ go build
student@ubuntu:~/student/btreedeletenode$ ./btreedeletenode
Before delete:
1
4
5
7
After delete:
1
5
7
student@ubuntu:~/student/btreedeletenode$ 
```
