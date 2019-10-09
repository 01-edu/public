## listmerge

### Instructions

Écrire une fonction `ListMerge` qui place les éléments d'une liste `l2` à la fin d'une autre liste `l1`.

-   Des nouveaux éléments ne doivent pas être créés!

### Fonction et structure attendues

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListMerge(l1 *List, l2 *List) {

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

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func main() {
	link := &piscine.List{}
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "a")
	piscine.ListPushBack(link, "b")
	piscine.ListPushBack(link, "c")
	piscine.ListPushBack(link, "d")
	fmt.Println("-----first List------")
	PrintList(link)

	piscine.ListPushBack(link2, "e")
	piscine.ListPushBack(link2, "f")
	piscine.ListPushBack(link2, "g")
	piscine.ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	PrintList(link2)

	fmt.Println("-----Merged List-----")
	piscine.ListMerge(link, link2)
	PrintList(link)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
a -> b -> c -> d -> e -> f -> g -> h -> <nil>
student@ubuntu:~/piscine-go/test$
```
