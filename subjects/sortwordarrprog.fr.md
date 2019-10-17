## sortwordarr

### Instructions

Écrire une fonction `SortWordArr` qui trie par ordre `ascii` (dans l'ordre croissant) un tableau de `string`.

### Fonction attendue

```go
func SortWordArr(array []string) {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	"fmt"
)

func main() {

	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}
```

Et son résultat :

```console
student@ubuntu:~/test$ go build
student@ubuntu:~/test$ ./test
[1 2 3 A B C a b c]
student@ubuntu:~/test$
```
