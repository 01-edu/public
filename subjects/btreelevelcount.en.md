## btreelevelcount

### Instructions

Write a function, `BTreeLevelCount`, that returns the number of levels of the binary tree (height of the tree)

### Expected function

```go
func BTreeLevelCount(root *TreeNode) int {

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
	fmt.Println(piscine.BTreeLevelCount(root))
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
3
student@ubuntu:~/[[ROOT]]/test$
```
