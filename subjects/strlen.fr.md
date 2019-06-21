## strlen

### Instructions

- Écrire une fonction qui compte le nombre de caractères d'une chaîne et qui retourne le nombre trouvé.

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
	piscine ".."
)

func main() {
	str := "Hello World!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
12
student@ubuntu:~/piscine/test$
```
