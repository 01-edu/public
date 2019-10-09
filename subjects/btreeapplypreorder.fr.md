## btreeapplypreorder

### Instructions

Écrire une fonction qui applique une fonction en pré-ordre ("preorder walk") à chaque élément de l'arbre.

### Fonction attendue

```go
func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {

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
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	piscine.BTreeApplyPreorder(root, fmt.Println)

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
1
7
5
student@ubuntu:~/piscine-go/test$
```
