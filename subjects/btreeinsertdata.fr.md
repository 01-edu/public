## btreeinsertdata

### Instructions

Écrire une fonction qui insère de la nouvelle donnée dans un arbre binaire en suivant les propriétés des arbres de recherche binaires.

Les nodes doivent être définies comme ci-dessous :

### Fonction attendue

```go
type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                 string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {

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
     fmt.Println(root.Left.Data)
     fmt.Println(root.Data)
     fmt.Println(root.Right.Left.Data)
     fmt.Println(root.Right.Data)

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/btreeinsertdata$ go build
student@ubuntu:~/piscine-go/btreeinsertdata$ ./btreeinsertdata
1
4
5
7
student@ubuntu:~/piscine-go/btreeinsertdata$
```
