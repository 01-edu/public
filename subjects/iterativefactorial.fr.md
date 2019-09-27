## iterativefactorial

### Instructions

Écrire une fonction **itérative** qui renvoie la factorielle d'un `int` passé en paramètre.

Les erreurs (valeurs non possibles ou overflows) renverront `0`.

### Fonction attendue

```go
func IterativeFactorial(nb int) int {

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
	fmt.Println(piscine.IterativeFactorial(arg))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
24
student@ubuntu:~/piscine-go/test$
```
