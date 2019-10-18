## divmod

### Instructions

-   Écrire une fonction qui aura le format ci-dessous.

### Fonction attendue

```go
func DivMod(a int, b int, div *int, mod *int) {

}
```

-   Cette fonction divisera les int **a** et **b**.
-   Le résultat de la division sera stocké dans l'int pointé par **div**.
-   Le reste de cette division sera stocké dans l'int pointé par **mod**.

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
    "fmt"
    piscine ".."
)

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	piscine.DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
6
1
student@ubuntu:~/piscine-go/test$
```
