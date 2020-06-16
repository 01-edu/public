## btreemax

### Instructions

Write a function, `BTreeMax`, that returns the node with the maximum value in the tree given by `root`.

### Expected function

```go
func BTreeMax(root *TreeNode) *TreeNode {

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
	max := piscine.BTreeMax(root)
	fmt.Println(max.Data)
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
7
student@ubuntu:~/[[ROOT]]/test$
```
