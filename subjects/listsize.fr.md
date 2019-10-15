## listsize

### Instructions

Écrire une fonction `ListSize` qui retourne le nombre d'éléments dans une liste chaînée `l`.

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

func ListSize(l *List) int {

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

func main() {
	link := &piscine.List{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "2")
	piscine.ListPushFront(link, "you")
	piscine.ListPushFront(link, "man")

	fmt.Println(piscine.ListSize(link))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
student@ubuntu:~/piscine-go/test$
```
