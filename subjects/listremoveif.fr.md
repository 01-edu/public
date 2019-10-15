## listremoveif

### Instructions

Écrire une fonction `ListRemoveIf` qui supprime tous les éléments qui sont égaux à la `data_ref` introduite dans l'argument de la fonction.

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

func ListRemoveIf(l *List, data_ref interface{}) {

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

	fmt.Println("----normal state----")
	piscine.ListPushBack(link2, 1)
	PrintList(link2)
	piscine.ListRemoveIf(link2, 1)
	fmt.Println("------answer-----")
	PrintList(link2)
	fmt.Println()

	fmt.Println("----normal state----")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "There")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "How")
	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)
	PrintList(link)

	piscine.ListRemoveIf(link, 1)
	fmt.Println("------answer-----")
	PrintList(link)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
----normal state----
1 -> <nil>
------answer-----
<nil>

----normal state----
1 -> Hello -> 1 -> There -> 1 -> 1 -> How -> 1 -> are -> you -> 1 -> <nil>
------answer-----
Hello -> There -> How -> are -> you -> <nil>
student@ubuntu:~/piscine-go/test$
```
