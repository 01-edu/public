## prinstr

### Instructions

-   Écrire une fonction qui affiche un à un les caractères d'une `string` à l'écran.

### Fonction attendue

```go
func PrintStr(str string) {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import piscine ".."

func main() {
	str := "Hello World!"
	piscine.PrintStr(str)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test | cat -e
Hello World!%
student@ubuntu:~/piscine-go/test$
```
