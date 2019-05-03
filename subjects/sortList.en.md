## sortList

### Instructions

Write the following function:

```go
func SortList (l *NodeList, cmp func(a,b int) bool) *NodeList{
	
}
```
- This function must sort the list given as a parameter, using the function cmp to select the order to apply, and returns a pointer to the first element of the sorted list.

- Duplications must remain.

- Input will always be consistent.

- You must use the `type NodeList`.

- Functions passed as cmp will always return a value different from 0 if a and b are in the right order, 0 otherwise.

- For example, the following function used as cmp will sort the list in ascending order :

```go
func ascending(a, b int) bool{
	if a <= b {
		return true
	} else {
		return false
	}
}
```