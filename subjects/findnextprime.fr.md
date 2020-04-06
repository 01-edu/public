## findnextprime

### Instructions

Écrire une fonction qui renvoie le premier nombre premier qui est égal ou supérieur à l'`int` passé en paramètre.

La fonction devra être optimisée pour éviter les time-outs avec le testeur.

### Fonction attendue

```go
func FindNextPrime(nb int) int {

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
	arg1 := 5
	arg2 := 4
	fmt.Println(piscine.FindNextPrime(arg1))
	fmt.Println(piscine.FindNextPrime(arg2))
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
5
5
student@ubuntu:~/[[ROOT]]/test$
```
