## strrev

### Instructions

-   Écrire une fonction qui inverse une chaîne de caractères(`string`).

-   Cette fonction **retournera** la chaîne de caractères s(`string`).

### Fonction attendue

```go
func StrRev(s string) string {

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
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
!dlroW olleH
student@ubuntu:~/piscine-go/test$
```
