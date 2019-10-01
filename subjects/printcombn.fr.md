## printcombn

### Instructions

- Écrire une fonction qui affiche toutes les combinaisons possibles de **n** chiffres différents en ordre croissant.

- n sera défini tel que: 0 < n < 10

ci-dessous vos références pour le **format d'affichage** attendu.

- (pour n = 1) '0, 1, 2, 3, ...8, 9'

- (pour n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

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
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
01, 02, 03, 04, 05, 06, 07, 08, 09, 12, 13, 14, 15, 16, 17, 18, 19, 23, 24, 25, 26, 27, 28, 29, 34, 35, 36, 37, 38, 39, 45, 46, 47, 48, 49, 56, 57, 58, 59, 67, 68, 69, 78, 79, 89
012, 013, 014, 015, 016, 017, 018, 019, 023, 024, 025, 026, 027, 028, 029, 034, 035, 036, 037, 038, 039, 045, 046, 047, 048, 049, 056, 057, 058, 059, 067, 068, 069, 078, 079, 089, 123, 124, 125, 126, 127, 128, 129, 134, 135, 136, 137, 138, 139, 145, 146, 147, 148, 149, 156, 157, 158, 159, 167, 168, 169, 178, 179, 189, 234, 235, 236, 237, 238, 239, 245, 246, 247, 248, 249, 256, 257, 258, 259, 267, 268, 269, 278, 279, 289, 345, 346, 347, 348, 349, 356, 357, 358, 359, 367, 368, 369, 378, 379, 389, 456, 457, 458, 459, 467, 468, 469, 478, 479, 489, 567, 568, 569, 578, 579, 589, 678, 679, 689, 789
student@ubuntu:~/piscine/test$
```
