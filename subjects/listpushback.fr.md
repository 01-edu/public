## listpushback

### Instructions

Écrire une fonction `ListPushBack` qui insère un nouvel élément `NodeL` à la fin de la liste `l` en utilisant la structure `List`.

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

func ListPushBack(l *List, data interface{}) {

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

	piscine.ListPushBack(link, "Hello")
	piscine.ListPushBack(link, "man")
	piscine.ListPushBack(link, "how are you")

	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello
man
how are you
student@ubuntu:~/piscine-go/test$
```
