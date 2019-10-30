## max

### Instructions

Écrire une fonction, `Max`, qui retourne la valeur maximum dans une slice d'entiers.

### Fonction attendue

```go
func Max(arr []int) int {

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
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(arrInt)
	fmt.Println(max)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
123
student@ubuntu:~/piscine-go/test$
```
