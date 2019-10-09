## convertbase

### Instructions

Écrire une fonction qui retourne la convertion d'un nombre `string` d'une baseFrom `string` à une baseTo `string`.

Seules des bases valides seront testées.

Les nombres négatifs ne seront pas testés.

### Fonction attendue

```go
func ConvertBase(nbr, baseFrom, baseTo string) string {

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
	result := piscine.ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
43
student@ubuntu:~/piscine-go/test$
```
