## listfind

### Instructions

Écrire une fonction `ListFind` qui retourne l'adresse de la première node dans la liste `l` qui est déterminée comme étant égale à `ref` par la fonction `CompStr`.

-   Pour cet exercice la fonction `CompStr` doit être utilisée.

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

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {

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
	piscine.ListPushBack(link, "hello1")
	piscine.ListPushBack(link, "hello2")
	piscine.ListPushBack(link, "hello3")

	found := piscine.ListFind(link, interface{}("hello2"), piscine.CompStr)

	fmt.Println(found)
	fmt.Println(*found)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
0xc42000a0a0
hello2
student@ubuntu:~/piscine-go/test$
```

### Note

-   L'addresse peut être différente à chaque exécution du programme.
