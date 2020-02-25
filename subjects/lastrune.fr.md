## lastrune

### Instructions

Écrire une fonction qui retourne la dernière `rune` d'une `string`.

### Fonction attendue

```go
func LastRune(s string) rune {

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
	z01.PrintRune(piscine.LastRune("Hello!"))
	z01.PrintRune(piscine.LastRune("Salut!"))
	z01.PrintRune(piscine.LastRune("Ola!"))
	z01.PrintRune('\n')
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
!!!
student@ubuntu:~/[[ROOT]]/test$
```
