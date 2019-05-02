## isprime

### Intructions

Écrire une fonction qui renvoie `true` si l'`int` passé en paramètre est un nombre premier. Autrement elle renvoie `false`.

### Fonction attendue

```go
func IsPrime(int nb) bool {

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
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
true
false
student@ubuntu:~/piscine/test$
```
