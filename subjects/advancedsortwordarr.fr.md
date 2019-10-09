## advancedsortwordarr

### Instructions

Écrire une fonction `AdvancedSortWordArr` qui trie un tableau de `string`, basé sur la fonction `f` passée en paramètre.

### Fonction attendue

```go
func AdvancedSortWordArr(array []string, f func(a, b string) int) {

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

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	piscine.AdvancedSortWordArr(result, piscine.Compare)

	fmt.Println(result)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/piscine-go/test$
```
