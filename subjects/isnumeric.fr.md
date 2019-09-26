## isnumeric

### Instructions

Écrire une fonction qui retourne `true` si la `string` passée en paramètre contient seulement des caractères numériques, et qui retourne `false` autrement.

### Fonction attendue

```go
func IsNumeric(str string) bool {

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
	fmt.Println(piscine.IsNumeric("010203"))
	fmt.Println(piscine.IsNumeric("01,02,03"))

}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
true
false
student@ubuntu:~/piscine-go/test$
```
