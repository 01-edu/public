## concatparams

### Instructions

Écrire une fonction qui prend les arguments en paramètres et les retournes dans une `string`.

Les arguments doivent être **séparés** par un `\n`.

### Fonction attendue

```go
func ConcatParams(args []string) string {

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
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello
how
are
you?
student@ubuntu:~/piscine-go/test$
```
