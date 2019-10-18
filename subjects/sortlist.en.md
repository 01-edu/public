## sortlist

### Instructions

Write a function that must:

-   Sort the list given as a parameter, using the function cmp to select the order to apply,

-   Return a pointer to the first element of the sorted list.

Duplications must remain.

Inputs will always be consistent.

The `type NodeList` must be used.

Functions passed as `cmp` will always return `true` if `a` and `b` are in the right order, otherwise it will return `false`.

### Expected function

```go
type Nodelist struct {
	Data int
	Next *Nodelist
}

func SortList (l *NodeList, cmp func(a,b int) bool) *NodeList{

}
```

-   For example, the following function used as `cmp` will sort the list in ascending order :

```go
func ascending(a, b int) bool{
	if a <= b {
		return true
	} else {
		return false
	}
}
```
