## capitalize

### Instructions

Écrire une fonction qui met en majuscule la première lettre de chaque mot et en minuscule les autres lettres du reste du mot d'une `string`.

Un mot est une suite de caractères **alphanumériques**.

### Fonction attendue

```go
func Capitalize(s string) string {

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
	fmt.Println(piscine.Capitalize("Hello! How are you? How+are+things+4you?"))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
Hello! How Are You? How+Are+Things+4you?
student@ubuntu:~/piscine-go/test$
```
