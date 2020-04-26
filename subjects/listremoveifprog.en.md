## listremoveifprog

### Instructions

Write a function `ListRemoveIf` that removes all elements that have a `Data` field equal to the `data_ref` in the argument of the function.

### Expected function and structure

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {

}
```
