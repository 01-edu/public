## activebits

### Instructions

Écrire une fonction, `ActiveBits`, qui retourne le nombre de `bits` actifs (`bits` ayant la valeur 1) dans la représentation binaire d'un nombre entier.

### Fonction attendue

```go
func ActiveBits(n int) uint {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	nbits := piscine.ActiveBits(7)
	fmt.Println(nbits)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
3
student@ubuntu:~/[[ROOT]]/test$
```
