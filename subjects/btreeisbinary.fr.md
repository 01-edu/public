## btreeisbinary

### Instructions

Écrire une fonction, `BTreeIsBinary`, qui retourne `true` seulement si l'arbre donné par `root` suit les propriétés des arbres de recherche binaires.

### Fonction attendue

```go
func BTreeIsBinary(root *TreeNode) bool {

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
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	fmt.Println(piscine.BTreeIsBinary(root))
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
true
student@ubuntu:~/[[ROOT]]/test$
```
