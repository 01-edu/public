## btreemax

### Instructions

Write a function, `BTreeMax`, that returns the node with the maximum value in the tree given by `root`.

### Expected function

```go
func BTreeMax(root *TreeNode) *TreeNode {

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
	max := piscine.BTreeMax(root)
	fmt.Println(max.Data)
}
```

And its output :

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
7
student@ubuntu:~/student/test$
```
