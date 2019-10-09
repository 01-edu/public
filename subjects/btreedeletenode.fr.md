## btreedeletenode

### Instructions

Écrire une fonction, `BTreeDeleteNode`, qui efface `node` d'un arbre donné par `root`.

L'arbre en résultant devra toujours suivre les règles des arbres de recherche binaires.

### Fonction attendue

```go
func BTreeDeleteNode(root, node *TreeNode) *TreeNode {

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
	node := piscine.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
	root = piscine.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	piscine.BTreeApplyInorder(root, fmt.Println)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Before delete:
1
4
5
7
After delete:
1
5
7
student@ubuntu:~/piscine-go/test$
```
