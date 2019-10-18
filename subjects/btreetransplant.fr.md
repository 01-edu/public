## btreetransplant

### Instructions

Afin de déplacer les sous-arbres dans l'arbre de recherche binaire, écrire une fonction, `BTreeTransplant`, qui remplace le sous-arbre commencé par `node` avec la node `rplc` dans l'arbre donné par `root`.

### Fonction attendue

```go
func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

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
	node := piscine.BTreeSearchItem(root, "1")
	replacement := &piscine.TreeNode{Data: "3"}
	root = piscine.BTreeTransplant(root, node, replacement)
	piscine.BTreeApplyInorder(root, fmt.Println)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
3
4
5
7
student@ubuntu:~/piscine-go/test$
```
