## btreeapplypreorder

### Instructions

Write a function that applies a function using a preorder walk to each element in the tree.

### Expected function

```go
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {

}
```

### Usage

Here is a possible program to test your function :

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
	piscine.BTreeApplyPreorder(root, fmt.Println)

}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
4
1
7
5
student@ubuntu:~/[[ROOT]]/test$
```
