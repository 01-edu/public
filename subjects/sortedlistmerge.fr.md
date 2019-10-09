## sortedlistmerge

### Instructions

Écrire une fonction `SortedListMerge` qui merge deux listes `n1` et `n2` en ordre ascendant.

-   Pendant les tests `n1` et `n2` seront déjà triées.

### Fonction et structure attendues

```go
func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {

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
	var link2 *piscine.NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(piscine.SortedListMerge(link2, link))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
-2 -> 3 -> 5 -> 7 -> 9 -> <nil>
student@ubuntu:~/piscine-go/test$
```
