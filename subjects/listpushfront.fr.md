## listpushfront

### Instructions

Écrire une fonction `ListPushFront` qui insère un nouvel élément `NodeL` au début de la liste `l` en utilisant la structure `List`.

### Fonction et structue attendues

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {

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

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "man")
	piscine.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
how are you man Hello
student@ubuntu:~/[[ROOT]]/test$
```
