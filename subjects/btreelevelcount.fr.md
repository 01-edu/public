## btreelevelcount

### Instructions

Écrire une fonction, `BTreeLevelCount`, qui retourne le nombre de niveaux de l'arbre binaire. (la hauteur de l'arbre)

### Fonction attendue

```go
func BTreeLevelCount(root *TreeNode) int {

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
	fmt.Println(piscine.BTreeLevelCount(root))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
3
student@ubuntu:~/piscine-go/test$
```
