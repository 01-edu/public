## listreverse

### Instructions

Écrire une fonction `ListReverse` qui inverse l'ordre des éléments d'une liste chaînée `l` donnée.

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

func ListReverse(l *List) {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, 2)
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, 4)

	piscine.ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
4
3
2
1
Tail <nil>
Head &{4 0xc42000a140}
student@ubuntu:~/[[ROOT]]/test$
```
