## split

### Instructions

Écrire une fonction qui sépare les mots d'une `string` et les met dans un tableau de `string`.

Les séparateurs sont les caractères de la `string charset` donnée en paramètre.

### Fonction attendue

```go
func Split(str, charset string) []string {

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
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(piscine.Split(str, "HA"))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
[Hello how are you?]
student@ubuntu:~/piscine-go/test$
```
