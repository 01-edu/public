## printnbr

### Instructions

Écrire une fonction qui affiche un `int` passé en paramètre.
Toutes les valeurs de type `int` doivent être affichables.
Vous ne pouvez pas convertir en `int64`.

### Fonction attendue

```go
func PrintNbr(n int) {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import piscine ".."

func main() {
	piscine.PrintNbr(-123)
	piscine.PrintNbr(0)
	piscine.PrintNbr(123)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
-1230123
student@ubuntu:~/[[ROOT]]/test$
```
