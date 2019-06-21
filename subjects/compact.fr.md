## compact

### Instructions

Écrire une fonction `Compact` qui prend un pointeur sur tableau comme paramètre et qui réécris sur les éléments qui pointent sur `nil`.

- Indice : Cette fonction existe in Ruby.

### Fonction attendue

```go
func Compact(ptr *[]string, length int) int {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import fmt

func main() {
	array := []string{"hello", " ", "there", " ", "bye"}

	ptr := &array
	fmt.Println(Compact(ptr, len(array)))
}
```

Et son résultat :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
3
student@ubuntu:~/piscine/test$
```
