## btreeapplypostorder

### Instructions

Écrire une fonction qui applique une fonction en post-ordre ("postorder walk") à chaque élément de l'arbre.

### Fonction attendue

```go
func BTreeApplyPostorder(root *piscine.TreeNode, f func(...interface{}) (int, error)) {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine "."
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	BTreeApplyPostorder(root, fmt.Println)

}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/btreeinsertdata$ go build
student@ubuntu:~/[[ROOT]]/btreeinsertdata$ ./btreeinsertdata
1
5
7
4
student@ubuntu:~/[[ROOT]]/btreeinsertdata$
```
