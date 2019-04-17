# btreeisbinary
## Instructions

Write a function, BTreeIsBinary, that returns true only if the tree given by root follows the binary search tree properties.

This function must have the following signature.

## Expected function

```go
func BTreeIsBinary(root *TreeNode) bool {
	
}

```

## Usage

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
	fmt.Println(student.BTreeIsBinary(root))
}
```

And its output :

```console
student@ubuntu:~/student/btreeisbinary$ go build
student@ubuntu:~/student/btreeisbinary$ ./btreeisbinary 
true
student@ubuntu:~/student/btreeisbinary$ 
```
