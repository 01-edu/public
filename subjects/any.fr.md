## any

### Instructions

Écrire une fonction `Any` qui retournes `true`, pour un tableau de `string`:

- si, lorsque ce tableau de `string` est passé à travers une fonction `f`, au moins un element retournes `true`.

### Fonction attendue

```go
func Any(f func(string) bool, arr []string) bool {
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
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "is", "4", "you"}

	result1 := piscine.Any(piscine.IsNumeric, tab1)
	result2 := piscine.Any(piscine.IsNumeric, tab2)

	fmt.Println(result1)
	fmt.Println(result2)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
false
true
student@ubuntu:~/piscine/test$
```
