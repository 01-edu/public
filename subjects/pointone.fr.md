## pointone

### Instructions

-   Écrire une fonction qui prend un **pointeur sur int** en argument et qui assignes à cet int la valeur 1.

### Fonction attendue

```go
func PointOne(n *int) {

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
    n := 0
    piscine.PointOne(&n)
    fmt.Println(n)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
student@ubuntu:~/piscine-go/test$
```
