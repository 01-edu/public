## fibonacci

### Instructions

Écrire une fonction **récursive** qui renvoie la valeur de la suite de fibonacci correspondant à l'index passé en paramètre.

La premiére valeur est à l'index `0`.

La suite débute ainsi: 0, 1, 1, 2, 3 etc...

Un index négatif renvoie `-1`.

`for` est **interdit** pour cet exercice.

### Fonction attendue

```go
package piscine

func Fibonacci(index int) int {

}
```

### Utilisation

Voici un éventuel `main.go` :

```go
package main

import (
        "fmt"
        piscine ".."
)

func main() {
	arg1 := 4
	fmt.Println(piscine.Fibonacci(arg1))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$go build
student@ubuntu:~/piscine-go/test$ ./test
3
student@ubuntu:~/piscine-go/test$
```
