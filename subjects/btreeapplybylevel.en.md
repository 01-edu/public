## btreeapplybylevel

### Instructions

Write a function, `BTreeApplyByLevel`, that applies the function given by `fn` to each node of the tree given by `root`.

### Expected function

```go
func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error))  {

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
	piscine.BTreeApplyByLevel(root, fmt.Println)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
1
7
5
student@ubuntu:~/piscine-go/test$
```
