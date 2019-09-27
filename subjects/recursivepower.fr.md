## recursivepower

### Instructions

Écrire une fonction **récursive** qui renvoie la puissance de deux `int` passés en paramètre.

Les puissances négatives renverront `0`. Les overflows **ne doivent pas** être gérés.

`for` est **interdit** pour cet exercice.

### Fonction attendue

```go
func RecursivePower(nb int, power int) int {

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
	fmt.Println(piscine.RecursivePower(arg1, arg2))
}
```

Et son résultat :

```console
student@ubuntu:~$ go build
student@ubuntu:~$ ./recursivepower
64
student@ubuntu:~$
```
