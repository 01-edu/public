## btreelevelcount

### Instructions

Write a function, BTreeLevelCount, that return the number of levels of the tree (height of the tree)

### Expected function

```go
func BTreeLevelCount(root *piscine.TreeNode) int {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
  "fmt"
  student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	fmt.Println(BTreeLevelCount(root))
}
```

And its output :

```console
student@ubuntu:~/student/btreesearchitem$ go build
student@ubuntu:~/student/btreesearchitem$ ./btreesearchitem
3
student@ubuntu:~/student/btreesearchitem$
```
