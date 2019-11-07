## strlen

### Instructions

-   Écrire une fonction qui compte le nombre de `runes` d'une `string` et qui retourne le nombre trouvé.

### Fonction attendue

```go
func StrLen(str string) int {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"
)

func main() {
	str := "Hello World!"
	nb := StrLen(str)
	fmt.Println(nb)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
12
student@ubuntu:~/piscine-go/test$
```
