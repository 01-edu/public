## merge-it

### Instructions

Given two binary trees and imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of new tree.

Write a function, `MergeTrees`, that returns merged tree .
Note: The merging process must start from the root nodes of both trees.

Example 1:

Input:

          1
         / \
        2   3

       [1,2,3]

          1
         / \
        2   3

       [1,2,3]

Merged Tree:

          2
         / \
        4   6

       [2,4,6]

### Expected function

```go
type TreeNodeM struct {
    Left    *TreeNodeM
    Val     int
    Right   *TreeNodeM
}


func MergeTrees(t1 *TreeNodeM, t2 *TreeNodeM) *TreeNodeM {

}
```

### Usage

Here is a possible program to test your function :

```go
package main

func main() {
  mergedTree := &TreeNodeM{}
  t1 := NewRandTree()
  t2 := NewRandTree()

  mergedTree = MergeTrees(t1, t2)
  printTree(mergedTree)
}
```

### Output

```console
$ go run .
[20, 2, 3, 4, 5, 6, 6, 7, 3, 5]
$
```
