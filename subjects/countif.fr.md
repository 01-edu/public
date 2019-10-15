## countif

### Instructions

Écrire une fonction `CountIf` qui retournes le nombre d'éléments d'un tableau de `string` pour lesquels la fonction `f` retourne `true`.

### Fonction attendue

```go
func CountIf(f func(string) bool, tab []string) int {

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
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This","1", "is", "4", "you"}
	answer1 := piscine.CountIf(piscine.IsNumeric, tab1)
	answer2 := piscine.CountIf(piscine.IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
0
2
student@ubuntu:~/piscine-go/test$
```
