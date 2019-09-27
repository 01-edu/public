## btreeprintroot

### Instructions

Écrire une fonction qui affiche la valeur de node `root` d'un arbre binaire.
Les nodes doivent être définies comme ci-dessous:

### Fonction attendue

```go
type TreeNode struct {
	left, right *TreeNode
	data        string
}

func PrintRoot(root *TreeNode){

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

func main() {
     //rootNode initialized with the value "who"
     //rootNode1 initialized with the value "are"
     //rootNode2 initialized with the value "you"
     printRoot(rootNode)
     printRoot(rootNode1)
     printRoot(rootNode2)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/printroot$ go build
student@ubuntu:~/piscine-go/printroot$ ./printroot
who
are
you
student@ubuntu:~/piscine-go/test$
```
