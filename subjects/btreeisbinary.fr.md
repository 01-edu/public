## btreeisbinary

### Instructions

Écrire une fonction, `BTreeIsBinary`, qui retourne `true` seulement si l'arbre donné par `root` suit les propriétés des arbres de recherche binaires.

### Fonction attendue

```go
func BTreeIsBinary(root *TreeNode) bool {

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
	fmt.Println(piscine.BTreeIsBinary(root))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
true
student@ubuntu:~/piscine-go/test$
```
