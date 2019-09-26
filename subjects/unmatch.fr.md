## unmatch

### Instructions

Écrire une fonction, `Unmatch`, qui retourne l'élément de la slice (arr) qui n'a pas de paire correspondante.

### Fonction attendue

```go
func Unmatch(arr []int) int {

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
	arr := []int{1, 2, 3, 1, 2, 3, 4}
	unmatch := piscine.Unmatch(arr)
	fmt.Println(unmatch)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
4
student@ubuntu:~/piscine-go/test$
```
