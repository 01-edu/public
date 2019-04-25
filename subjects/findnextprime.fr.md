## findnextprime

### Intructions

Écrire une fonction qui renvoie le premier nombre premier qui est égal ou supérieur à l'`int` passé en paramètre.

### Fonction attendue

```go
func FindNextPrime(int nb) int {

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
	arg1 := 5
	arg2 := 4
	fmt.Println(piscine.FindNextPrime(arg1))
	fmt.Println(piscine.FindNextPrime(arg2))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
5
5
student@ubuntu:~/piscine/test$
```
