## islower

### Instructions

Écrire une fonction qui retourne `true` si la `string` passée en paramètre contient seulement des caractères minuscules, et qui retourne `false` autrement.

### Fonction attendue

```go
func IsLower(str string) bool {

}
```

### Utilisation

Voici un éventuel programmes pour tester votre fonction :

```go
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	fmt.Println(piscine.IsLower("hello"))
	fmt.Println(piscine.IsLower("hello!"))

}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
true
false
student@ubuntu:~/[[ROOT]]/test$
```
