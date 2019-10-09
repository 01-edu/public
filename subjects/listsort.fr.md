## listsort

### Instructions

Écrire une fonction `ListSort` qui trie les nodes d'une liste chaînée par ordre croissant.

-   La structure `NodeI` sera la seule utilisée.

### Fonction et structure attendues

```go
type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"

	piscine ".."
)

func PrintList(l *piscine.NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *piscine.NodeI, data int) *piscine.NodeI {
	n := &piscine.NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *piscine.NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(piscine.ListSort(link))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1 -> 2 -> 3 -> 4 -> 5 -> <nil>
student@ubuntu:~/piscine-go/test$
```
