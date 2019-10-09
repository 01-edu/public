## swap

### Instructions

-   Écrire une fonction qui échange les contenus de deux **pointeurs sur entier** (`*int`).

### Fonction attendue

```go
func Swap(a *int, b *int) {

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
	b := 1
	piscine.Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
1
0
student@ubuntu:~/piscine-go/test$
```
