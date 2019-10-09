# options

## Instructions

Écrire un programme qui prend un nombre indéfini d'arguments qui peuvent être considérés comme des `options` et qui affiche sur la sortie standard une représentation de ces `options` comme groupes de `bytes`(octets) suivi d'un retour à la ligne (`'\n'`).

-   Une `option` est un argument qui commence avec un `-` et qui peux avoir de multiples caractères qui peuvent être :
    -abcdefghijklmnopqrstuvwxyz

-   Toutes les `options` sont stockées dans un seul `int` et chaque `option` représente un bit de cet `int`, et doit être stocké comme ci-dessous :

        	- 00000000 00000000 00000000 00000000
        	- ******zy xwvutsrq ponmlkji hgfedcba

-   L'éxécution du programme sans argument ou avec l'option `-h` activée doit afficher toutes les `options` valides sur la sortie standard, comme montré dans un des exemples ci-dessous.

-   Une mauvaise `option` doit afficher `Invalid Option` suivi d'un retour à la ligne.

## Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test | cat -e
options: abcdefghijklmnopqrstuvwxyz$
student@ubuntu:~/piscine-go/test$ ./test -abc -ijk | cat -e
00000000 00000000 00000111 00000111$
student@ubuntu:~/piscine-go/test$ ./test -z | cat -e
00000010 00000000 00000000 00000000$
student@ubuntu:~/piscine-go/test$ ./test -abc -hijk | cat -e
options: abcdefghijklmnopqrstuvwxyz$
student@ubuntu:~/piscine-go/test$ ./test -% | cat -e
Invalid Option$
student@ubuntu:~/piscine-go/test$
```
