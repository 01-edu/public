## btreemin

### Instructions

Écrire une fonction, `BTreeMin`, qui retourne la node avec la valeur minimum de l'arbre donné par `root`.

### Fonction attendue

```go
func BTreeMin(root *TreeNode) *TreeNode {

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
	min := piscine.BTreeMin(root)
	fmt.Println(min.Data)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
1
student@ubuntu:~/[[ROOT]]/test$
```
