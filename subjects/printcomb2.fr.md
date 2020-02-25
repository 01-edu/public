## printcomb2

### Instructions

Écrire une fonction qui affiche sur une seule ligne dans l'ordre croissant toutes les combinaisons possibles de deux nombres différents à deux chiffres.

Les combinaisons sont séparées par une virgule et un espace.

### Fonction attendue

```go
func PrintComb2() {

}
```

### Utilisation

Voici un éventuel programme pour tester votre fonction :

```go
package main

import piscine ".."

func main() {
	piscine.PrintComb2()
}
```

Voici la sortie tronquée :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test | cat -e
00 01, 00 02, 00 03, ..., 00 98, 00 99, 01 02, 01 03, ..., 97 98, 97 99, 98 99$
student@ubuntu:~/[[ROOT]]/test$
```
