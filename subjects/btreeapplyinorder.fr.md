## btreeapplyinorder

### Instructions

Écrire une fonction qui applique une fonction en ordre (in order) à chaque élément de l'arbre (voir les "in order tree walks").

### Fonction attendue

```go
func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {

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
	piscine.BTreeApplyInorder(root, fmt.Println)

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
4
5
7
student@ubuntu:~/piscine-go/test$
```
