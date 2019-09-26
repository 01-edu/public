## abort

### Instructions

Écrire une fonction qui retournes la médiane de 5 arguments.

### Fonction attendue

```go
func Abort(a, b, c, d, e int) int {

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
	middle := piscine.Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
5
student@ubuntu:~/piscine-go/test$
```
