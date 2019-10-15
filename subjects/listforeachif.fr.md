## listforeachif

### Instructions

Écrire une fonction `ListForEachIf` qui applique un fonction donnée en argument à la data contenue dans certaines des nodes d'une liste `l`.

-   Cette fonction reçoit deux fonctions:

    -   `f` est la fonction qui est appliqué à la node.

    -   `cond` est une fonction qui retourne un `boolean` et qui sera utilisée pour déterminer si la fonction`f` doit être appliquée à la node.

-   La fonction donnée en argument doit avoir un pointeur `*NodeL` comme argument.

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

func IsPositive_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNegative_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	case string, rune:
		return false
	}
	return false
}

func IsNotNumeric_node(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	case string, rune:
		return true
	}
	return true
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	piscine ".."
	"fmt"
)

func PrintElem(node *piscine.NodeL) {
	fmt.Println(node.Data)
}

func StringToInt(node *piscine.NodeL) {
	node.Data = 2
}

func PrintList(l *piscine.List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, "->")
		it = it.Next
	}
	fmt.Print("nil","\n")
}

func main() {
	link := &piscine.List{}

	piscine.ListPushBack(link, 1)
	piscine.ListPushBack(link, "hello")
	piscine.ListPushBack(link, 3)
	piscine.ListPushBack(link, "there")
	piscine.ListPushBack(link, 23)
	piscine.ListPushBack(link, "!")
	piscine.ListPushBack(link, 54)

	PrintList(link)

	fmt.Println()
	fmt.Println("--------function applied--------")
	piscine.ListForEachIf(link, PrintElem, piscine.IsPositive_node)

	piscine.ListForEachIf(link, StringToInt, piscine.IsNotNumeric_node)

	fmt.Println("--------function applied--------")
	PrintList(link)

	fmt.Println()
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1->hello->3->there->23->!->54->nil
--------function applied--------
1
3
23
54
--------function applied--------
1->2->3->2->23->2->54->nil
student@ubuntu:~/piscine-go/test$
```
