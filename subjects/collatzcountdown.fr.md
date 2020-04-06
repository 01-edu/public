## collatzcountdown

### Instructions

Écrire une fonction, `CollatzCountdown`, qui retournes le nombre d'étapes nécéssaires pour atteindre 1 en utilisant le comptage de collatz.

- Elle doit renvoyer `-1` si `start` est égal à 0 ou négatif.

### Fonction attendue

```go
func CollatzCountdown(start int) int {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

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
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
10
student@ubuntu:~/[[ROOT]]/test$
```
