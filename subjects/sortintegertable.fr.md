## sortintegertable

### Instructions

- Écrire une fonction qui trie un tableau d'`int` (entier) par ordre croissant.

### Fonction attendue

```go
func SortIntegerTable(table []int) {

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
	s := []int{5,4,3,2,1,0}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[0,1,2,3,4,5]
student@ubuntu:~/piscine-go/test$
```
