## recursivefactorial

### Instructions

Écrire une fonction **récursive** qui renvoie la factorielle d'un `int` passé en paramètre.

Les erreurs (valeurs non possibles ou overflows) renverront `0`.

`for` est **interdit** pour cet exercice.

### Fonction attendue

```go
func RecursiveFactorial(nb int) int {

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
	arg := 4
	fmt.Println(piscine.RecursiveFactorial(arg))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
24
student@ubuntu:~/piscine-go/test$
```
