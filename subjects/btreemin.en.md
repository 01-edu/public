## btreemin

### Instructions

Write a function, `BTreeMin`, that returns the node with the minimum value in the tree given by root

### Expected function

```go
func BTreeMin(root *TreeNode) *TreeNode {

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
	min := piscine.BTreeMin(root)
	fmt.Println(min.Data)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
student@ubuntu:~/piscine-go/test$
```
