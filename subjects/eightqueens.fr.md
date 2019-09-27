## eightqueens

### Instructions

Écrire une [fonction](TODO-LINK) qui affiche toutes les solutions du [problème des huit dames](https://en.wikipedia.org/wiki/Eight_queens_puzzle).

La récursion doit être utilisée pour résoudre ce problème.

L'affichage sera quelque chose comme ça :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
15863724
16837425
17468253
...
```

Chaque solution sera sur une ligne unique.
L'index du placement d'une reine commence à 1.
Elle se lit de gauche à droite et chaque chiffre est la position pour chacune des colonnes.
Les solutions seront affichées dans l'ordre croissant.

### Fonction attendue

```go
package main

func EightQueens() {

}
```
