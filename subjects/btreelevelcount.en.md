## btreelevelcount

### Instructions

Write a function, `BTreeLevelCount`, that returns the number of levels of the binary tree (height of the tree)

### Expected function

```go
func BTreeLevelCount(root *TreeNode) int {

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
	fmt.Println(piscine.BTreeLevelCount(root))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
3
student@ubuntu:~/piscine-go/test$
```
