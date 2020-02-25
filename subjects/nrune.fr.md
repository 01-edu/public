## nrune

### Instructions

Écrire une fonction qui retourne la énième `rune` d'une `string`.

### Fonction attendue

```go
func NRune(s string, n int) rune {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import (
	"github.com/01-edu/z01"
	piscine ".."
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
la!
student@ubuntu:~/[[ROOT]]/test$
```
