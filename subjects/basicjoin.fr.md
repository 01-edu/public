## basicjoin

### Instructions

Écrire une fonction qui retourne la concaténation de toutes les `string` d'un slice de `string` passées en paramètres.

### Fonction attendue

```go
func BasicJoin(strs []string) string {

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
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.BasicJoin(toConcat))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello! How are you?
student@ubuntu:~/piscine-go/test$
```
