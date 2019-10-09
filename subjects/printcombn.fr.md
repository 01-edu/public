## printcombn

### Instructions

-   Écrire une fonction qui affiche toutes les combinaisons possibles de **n** chiffres différents en ordre croissant.

-   n sera défini tel que: 0 < n < 10

ci-dessous vos références pour le **format d'affichage** attendu.

-   (pour n = 1) '0, 1, 2, 3, ...8, 9'

-   (pour n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

### Fonction attendue

```go
func PrintCombN(n int) {

}
```

### Utilisation

Voici un éventuel [programme](TODO-LINK) pour tester votre fonction :

```go
package main

import piscine ".."

func main() {
	piscine.PrintCombN(1)
	piscine.PrintCombN(2)
	piscine.PrintCombN(3)
}
```

Et son résultat :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
012, 013, 014, 015, 016, 017, 018, ... 679, 689, 789
012345678, 012345679, ..., 123456789
student@ubuntu:~/piscine-go/test$
```
