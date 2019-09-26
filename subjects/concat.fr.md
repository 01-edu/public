## concat

### Instructions

Écrire une fonction qui retourne la concaténation de deux `string` passées en paramètres.

### Fonction attendue

```go
func Concat(str1 string, str2 string) string {

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
	fmt.Println(piscine.Concat("Hello!", " How are you?"))

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello! How are you?
student@ubuntu:~/piscine-go/test$
```
