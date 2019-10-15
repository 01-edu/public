## listforeach

### Instructions

Écrire une fonction `ListForEach` qui applique un fonction donnée en argument à la data contenue dans chacune des nodes d'une liste `l`.

-   La fonction donnée en argument doit avoir un pointeur comme argument: `l *List`

-   Copier les fonctions `Add2_node` et `Subtract3_node` dans le même fichier où la fonction `ListForEach` est définie.

### Fonction et struture attendues

```go
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListForEach(l *List, f func(*NodeL)) {
}

func Add2_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
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

	piscine.ListPushBack(link, "1")
	piscine.ListPushBack(link, "2")
	piscine.ListPushBack(link, "3")
	piscine.ListPushBack(link, "5")

	piscine.ListForEach(link, piscine.Add2_node)

	it := link.Head
	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12
22
32
52
student@ubuntu:~/piscine-go/test$
```
