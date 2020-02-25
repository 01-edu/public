## isnegative

### Instructions

Écrire une fonction qui affiche `'T'` (true) sur une seule ligne si l'`int` passé en paramètre est négatif, sinon elle affiche `'F'` (false).

### Fonction attendue

```go
func IsNegative(nb int) {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import piscine ".."

func main() {
	piscine.IsNegative(1)
	piscine.IsNegative(0)
	piscine.IsNegative(-1)
}
```

Et son résultat :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test
F
F
T
student@ubuntu:~/[[ROOT]]/test$
```
