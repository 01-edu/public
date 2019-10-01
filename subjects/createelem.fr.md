## createelem

### Instructions

Écrire une fonction `CreateElem` qui crée un nouvel élément de type `Node`.

### Fonction attendue et structure

```go
type Node struct {
	Data interface{}
}

func CreateElem(n *Node, value int) {

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
	n := &node{}
	n.CreateElem(1234)
	fmt.Println(n)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
&{1234}
student@ubuntu:~/piscine-go/test$
```
