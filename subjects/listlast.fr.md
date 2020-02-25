## listlast

### Instructions

Écrire une fonction `ListLast` qui retourne le dernier élément d'une liste chaînée `l`.

### Fonction et structure attendue

```go
type Node struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListLast(l *List) interface{} {

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
	link2 := &piscine.List{}

	piscine.ListPushBack(link, "three")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "1")

	fmt.Println(piscine.ListLast(link))
	fmt.Println(piscine.ListLast(link2))
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
1
<nil>
student@ubuntu:~/[[ROOT]]/test$
```
