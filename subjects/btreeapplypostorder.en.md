## btreeapplypostorder

### Instructions

Write a function that applies a function using a postorder walk to each element in the tree.

### Expected function

```go
func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {

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
	piscine.BTreeApplyPostorder(root, fmt.Println)

}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
5
7
4
student@ubuntu:~/piscine-go/test$
```
