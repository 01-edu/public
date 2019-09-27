## splitwhitespaces

### Instructions

Écrire une fonction qui sépare les mots d'une `string` et les met dans un tableau de `string`.

Les séparateurs sont les espaces, les tabulations et les retours à la ligne.

### Fonction attendue

```go
func SplitWhiteSpaces(str string) []string {

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
	str := "Hello how are you?"
	fmt.Println(piscine.SplitWhiteSpaces(str))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[Hello how are you?]
student@ubuntu:~/piscine-go/test$
```
