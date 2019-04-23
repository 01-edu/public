## btreeapplybylevel
### Instructions

Write a function, BTreeApplyByLevel, that applies the function given by fn to each node of the tree given by root.

This function must have the following signature.

### Expected function

```go
func BTreeApplyByLevel(root *TreeNode, fn interface{})  {
	
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
	student.BTreeApplyByLevel(root, fmt.Println)
}
```

And its output :

```console
student@ubuntu:~/student/btreeapplybylevel$ go build
student@ubuntu:~/student/btreeapplybylevel$ ./btreeapplybylevel
4
1
7
5
student@ubuntu:~/student/btreeapplybylevel$ 
```
