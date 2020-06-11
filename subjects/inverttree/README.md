## invert tree

### Instructions

Write a function that takes tree and inverts(flips) and returns it.

### Expected function and structure

```go
type TNode struct {
    Val     int
    Left    *TNode
    Right   *TNode
}

func InvertTree(root *TNode) *TNode {

}
```

Example:

```shell
Input:
      7
    /   \
   5     10
  / \    / \
 3   6  9  13

Output:
        7
     /    \
   10      5
  / \     / \
13   9   6   3
```
