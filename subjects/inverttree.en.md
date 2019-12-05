## invert tree

### Instructions:
Write a function that takes tree and inverts(flips) and returns it.
```
type Node struct {
    Val     int
    Left    *Node
    Right   *Node
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
func InvertTree(root *Node) *Node {
}
```
