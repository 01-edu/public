## btreeprintroot

### Instructions

Write a function to print the value of the root node of a binary tree.
The nodes must be defined as follows:

### Expected function

```go
type TreeNode struct {
	left, right *TreeNode
	data        string
}

func PrintRoot(root *TreeNode){

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

func main() {
     //rootNode initialized with the value "who"
     //rootNode1 initialized with the value "are"
     //rootNode2 initialized with the value "you"
     printRoot(rootNode)
     printRoot(rootNode1)
     printRoot(rootNode2)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/printroot$ go build
student@ubuntu:~/piscine-go/printroot$ ./printroot
who
are
you
student@ubuntu:~/piscine-go/test$
```
