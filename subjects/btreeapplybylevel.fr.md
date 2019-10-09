## btreeapplybylevel

### Instructions

Écrire une fonction, `BTreeApplyByLevel`, qui applique la fonction donnée par `fn` à chacune des nodes de l'arbre donné par `root`.

### Fonction attendue

```go
func BTreeApplyByLevel(root *TreeNode, fn interface{})  {

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
	piscine.BTreeApplyByLevel(root, fmt.Println)
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
