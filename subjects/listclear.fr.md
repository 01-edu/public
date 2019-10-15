## listclear

### Instructions

Écrire une fonction `ListClear` qui efface toutes les `nodes` d'une liste chaînée `l`.

-   Indice: assigner le pointeur de la liste à `nil`.

### Fonction et structure attendues

```go
func ListClear(l *List) {

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

type List = piscine.List
type Node = piscine.NodeL

func PrintList(l *List) {
	link := l.Head
	for link != nil {
		fmt.Print(link.Data, " -> ")
		link = link.Next
	}
	fmt.Println(nil)
}

func main() {
	link := &List{}

	piscine.ListPushBack(link, "I")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "something")
	piscine.ListPushBack(link, 2)

	fmt.Println("------list------")
	PrintList(link)
	piscine.ListClear(link)
	fmt.Println("------updated list------")
	PrintList(link)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
------list------
I -> 1 -> something -> 2 -> <nil>
------updated list------
<nil>
student@ubuntu:~/piscine-go/test$
```
