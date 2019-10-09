## makerange

### Instructions

Écrire une fonction qui prend un `int` minimum et un `int` maximum comme paramètres. Cette fonction retournes une slice d'`int` avec toutes les valeurs comprises entre le minimum et le maximum.

Le minimum est inclus, le maximum est exclu.

Si le minimum est supérieur ou égal au maximum, une slice `nil` est retournée.

`append` n'est pas autorisé pour cet exercice.

### Fonction attendue

```go
func MakeRange(min, max int) []int {

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
	fmt.Println(piscine.MakeRange(5, 10))
	fmt.Println(piscine.MakeRange(10, 5))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[5 6 7 8 9]
[]
student@ubuntu:~/piscine-go/test$
```
