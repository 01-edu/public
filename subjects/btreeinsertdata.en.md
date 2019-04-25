## btreeinsertdata

### Instructions

Write a function that inserts new data in a binary search tree
following the properties of binary search trees.
The nodes must be defined as follows:

### Expected function

```go
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

func main() {
     root := &TreeNode{data: "4"}
     BTreeInsertData(root, "1")
     BTreeInsertData(root, "7")
     BTreeInsertData(root, "5")
     fmt.Println(root.left.data)
     fmt.Println(root.data)
     fmt.Println(root.right.left.data)
     fmt.Println(root.right.data)

}
```

And its output :

```console
student@ubuntu:~/piscine/btreeinsertdata$ go build
student@ubuntu:~/piscine/btreeinsertdata$ ./btreeinsertdata
1
4
5
7
student@ubuntu:~/piscine/btreeinsertdata$
```
