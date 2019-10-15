## issorted

### Instructions

Écrire une fonction `IsSorted` qui retourne `true` si la slice d'`int` est triée, et qui retourne `false` autrement.

La fonction passée en paramètre retourne un `int` positive si `a` (le premier argument) est supérieur à `b`(le deuxième argument), elle retourne `0` si ils sont égaux et elle retourne un `int` négatif autrement.

Pour faire vos tests, vous devez coder votre propre fonction `f`.

### Fonction attendue

```go
func IsSorted(f func(a, b int) int, tab []int) bool {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction (sans `f`) :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	tab1 := []int{0, 1, 2, 3, 4, 5}
	tab2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, tab1)
	result2 := piscine.IsSorted(f, tab2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
true
false
student@ubuntu:~/piscine-go/test$
```
