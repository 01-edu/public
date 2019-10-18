## ultimatepointone

### Instructions

-   Écrire une fonction qui prend un **pointeur sur pointeur sur pointeur sur int** en argument et qui assignes à cet int la valeur 1.

### Fonction attendue

```go
func UltimatePointOne(n ***int) {

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
	a := 0
	b := &a
	n := &b
	piscine.UltimatePointOne(&n)
	fmt.Println(a)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
student@ubuntu:~/piscine-go/test$
```
