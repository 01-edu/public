## atoibase

### Instructions

Écrire une fonction qui prend un nombre `string` et sa base `string` en paramètres et retourne sa conversion en `int`.

Si la base n'est pas valide elle retourne `0`:

Règles de validité d'une base :

-   Une base doit contenir au moins 2 caractères.
-   Chaque caractère d'une base doit être unique.
-   Une base ne doit pas contenir les caractères `+` ou `-`.

Seuls des nombres en `string` valides seront testés.

La fonction **ne doit pas** gérer les nombres négatifs.

### Fonction attendue

```go
func AtoiBase(s string, base string) int {

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
	fmt.Println(piscine.AtoiBase("125", "0123456789"))
	fmt.Println(piscine.AtoiBase("1111101", "01"))
	fmt.Println(piscine.AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(piscine.AtoiBase("uoi", "choumi"))
	fmt.Println(piscine.AtoiBase("bbbbbab", "-ab"))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
125
125
125
125
0
student@ubuntu:~/piscine-go/test$
```
