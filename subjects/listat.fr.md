## listat

### Instructions

Écrire une fonction `ListAt` qui prend un pointeur sur la liste `l` et un `int pos` comme paramètres. Cette fonction devra afficher la `NodeL` à la position `pos` de la liste chaînée `l`.

-   En cas d'erreur la fonction affichera `nil`.

### Fonction et structure attendues

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}


func ListAt(l *NodeL, pos int) *NodeL{

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

	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, "how are")
	piscine.ListPushBack(link, "you")
	piscine.ListPushBack(link, 1)

	fmt.Println(piscine.ListAt(link.Head, 3).Data)
	fmt.Println(piscine.ListAt(link.Head, 1).Data)
	fmt.Println(piscine.ListAt(link.Head, 7))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
how are
<nil>
student@ubuntu:~/piscine-go/test$
```
