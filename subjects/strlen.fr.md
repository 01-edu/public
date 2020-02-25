## strlen

### Instructions

- Écrire une fonction qui compte le nombre de `runes` d'une `string` et qui retourne le nombre trouvé.

### Fonction attendue

```go
func StrLen(str string) int {

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
	str := "Hello World!"
	nb := piscine.StrLen(str)
	fmt.Println(nb)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
12
student@ubuntu:~/[[ROOT]]/test$
```
