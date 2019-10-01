## iterativepower

### Instructions

Écrire une fonction **itérative** qui renvoie la puissance de deux `int` passés en paramètre.

Les puissances négatives renverront `0`. Les overflows **ne doivent pas** être gérés.

### Fonction attendue

```go
func IterativePower(nb int, power int) int {

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
	arg1 := 4
	arg2 := 3
	fmt.Println(piscine.IterativePower(arg1, arg2))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
64
student@ubuntu:~/piscine-go/test$
```
