## ultimatedivmod

### Instructions

- Écrire une fonction qui aura le format ci-dessous.

### Fonction attendue

```go
func UltimateDivMod(a *int, b *int) {

}
```

- Cette fonction divisera les int pointés par **a** et **b**.
- Le résultat de la division sera stocké dans l'int pointé par **a**.
- Le reste de cette division sera stocké dans l'int pointé par **b**.

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
	piscine.UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
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
