## index

### Instructions

Écrire une fonction qui se comporte comme la fonction [`Index`](https://golang.org/pkg/strings/#Index).

### Fonction attendue

```go
func Index(s string, toFind string) int {

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
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
2
1
-1
student@ubuntu:~/piscine-go/test$
```
