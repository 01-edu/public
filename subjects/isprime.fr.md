## isprime

### Instructions

Écrire une fonction qui renvoie `true` si l'`int` passé en paramètre est un nombre premier. Autrement elle renvoie `false`.

La fonction devra être optimisée pour éviter les time-outs avec le testeur.

### Fonction attendue

```go
func IsPrime(nb int) bool {

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
	fmt.Println(piscine.IsPrime(5))
	fmt.Println(piscine.IsPrime(4))
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
