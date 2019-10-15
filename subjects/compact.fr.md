## compact

### Instructions

Écrire une fonction `Compact` qui prend un pointeur sur slice de `string` comme paramètre.
Cette fonction doit:

-   Retourner le nombre d'éléments avec des valeurs non-`nil`

-   Comprimer, c.à.d., effacer les éléments qui ont une valeur `nil` dans la slice.

### Fonction attendue

```go
func Compact(ptr *[]string) int {

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

const N = 6

func main() {
	arr := make([]string, N)
	arr[0] = "a"
	arr[2] = "b"
	arr[4] = "c"

	for _, v := range arr {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", piscine.Compact(&arr))

	for _, v := range arr {
		fmt.Println(v)
	}
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
a

b

c

Size after compacting: 3
a
b
c
student@ubuntu:~/piscine-go/test$
```
