## btreeisbinary

### Instructions

Write a function, `BTreeIsBinary`, that returns `true` only if the tree given by `root` follows the binary search tree properties.

### Expected function

```go
func BTreeIsBinary(root *TreeNode) bool {

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
	fmt.Println(piscine.BTreeIsBinary(root))
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
true
student@ubuntu:~/piscine-go/test$
```
