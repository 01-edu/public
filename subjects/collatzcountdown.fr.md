## collatzcountdown

### Instructions

Écrire une fonction, `CollatzCountdown`, qui retournes le nombre d'étapes nécéssaires pour atteindre 1 en utilisant le comptage de collatz.

-   Elle doit renvoyer `-1` si `start` est égal à 0 ou négatif.

### Fonction attendue

```go
func CollatzCountdown(start int) int {

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
	steps := piscine.CollatzCountdown(12)
	fmt.Println(steps)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
10
student@ubuntu:~/piscine-go/test$
```
