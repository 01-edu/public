## sortlistinsert

### Instructions

Écrire une fonction `SortListInsert` qui insère `data_ref` dans la liste chaînée `l`
tout en gardant cette liste triée par ordre croissant.

-   Pendant les tests la liste passée en argument sera déjà triée.

### Fonction et structure attendues

```go
func SortListInsert(l *NodeI, data_ref int) *NodeI{

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

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = piscine.SortListInsert(link, -2)
	link = piscine.SortListInsert(link, 2)
	PrintList(link)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1 -> 4 -> 9 -> <nil>
-2 -> 1 -> 2 -> 4 -> 9 -> <nil>
student@ubuntu:~/piscine-go/test$
```
