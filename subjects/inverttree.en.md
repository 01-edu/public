## invert tree

### Instructions:
Write a function that takes tree and inverts(flips) and returns it.
```
type TNode struct {
    Val     int
    Left    *TNode
    Right   *TNode
}
```
Example:
```
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
Expected function:
```
func InvertTree(root *TNode) *TNode {
}
```
