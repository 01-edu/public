## btreemax

### Instructions

Écrire une fonction, `BTreeMax`, qui retourne la node avec la valeur maximum de l'arbre donné par `root`.

### Fonction attendue

```go
func BTreeMax(root *TreeNode) *TreeNode {

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
	max := piscine.BTreeMax(root)
	fmt.Println(max.Data)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
7
student@ubuntu:~/piscine-go/test$
```
