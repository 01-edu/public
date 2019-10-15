## putnbrbase

### Instructions

Écrire une fonction qui affiche un `int` dans une base en `string` passés en paramètres.

Si la base n'est pas valide, la fonction affiche `NV` (Not Valid):

Règles de validité d'une base :

-   Une base doit contenir au moins 2 caractères.
-   Chaque caractère d'une base doit être unique.
-   Une base ne doit pas contenir les caractères `+` ou `-`.

La fonction doit gérer les nombres négatifs (comme montré sur l'exemple).

### Fonction attendue

```go
func PrintNbrBase(nbr int, base string) () {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"
	"github.com/01-edu/z01"
	piscine ".."
)

func main() {
	piscine.PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	piscine.PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
125
-1111101
7D
-uoi
NV
student@ubuntu:~/piscine-go/test$
```
