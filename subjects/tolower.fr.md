## tolower

### Instructions

Écrire une fonction qui met en minuscule chaque lettre d'une `string`.

### Fonction attendue

```go
func ToLower(s string) string {

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
	fmt.Println(piscine.ToLower("Hello! How are you?"))
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
hello! how are you?
student@ubuntu:~/[[ROOT]]/test$
```
