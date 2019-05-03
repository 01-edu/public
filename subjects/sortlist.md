## sortlist

### Instructions

Write the following functions and its struture :

```go
type Node struct {
	data int
	next *node
}

func SortList(l *node, func cmp(a,b int) bool) *node{

}
```
This function must sort the list given as a parameter, using the function `cmp` to select the order to apply, and returns a pointer to the first element of the sorted list.

- Duplications must remain.

- Inputs will always be consistet.

- You must use the type `Node`

- Functions passed as `cmp` will always return a boolean. If `a` and `b` are in the rigth order it returns `true`, otherwise it returns `false`.

- For example, the following function used as cmp will sort the list in ascending order :

```go
func ascending(a, b int) {
	if a <= b {
		return true
	} else {
		return false
	}
}
```