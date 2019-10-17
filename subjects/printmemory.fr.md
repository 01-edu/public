## printmemory

### Instructions

Écrire une fonction qui prend `(arr [10]int)`, et qui affiche la mémoire comme dans l'exemple.

### Fonction attendue

```go
func PrintMemory(arr [10]int) {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import (
	piscine ".."
)

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	piscine.PrintMemory(arr)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test | cat -e
6800 0000 6500 0000 6c00 0000 6c00 0000 $
6f00 0000 1000 0000 1500 0000 2a00 0000 $
0000 0000 0000 0000 $
hello..*..$
student@ubuntu:~/piscine-go/test$
```
