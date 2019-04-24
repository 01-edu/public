## btreeinsertdata

### Instructions

Write a function that applies a function using a preorder walk to each element in the tree

### Expected function

```go
func BTreeApplyPreorder(root *piscine.TreeNode, f func(...interface{}) (int, error)) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
       "fmt"
       piscine "."
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	BTreeApplyPreorder(root, fmt.Println)

}
```

And its output :

```console
student@ubuntu:~/piscine/btreeinsertdata$ go build
student@ubuntu:~/piscine/btreeinsertdata$ ./btreeinsertdata
4
1
7
5
student@ubuntu:~/piscine/btreeinsertdata$
```
