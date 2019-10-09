## firstword

### Instructions

Écrire un programme qui prend une `string` et qui affiche son premier mot, suivi d'un retour à la ligne (`'\n'`).

-   Un mot est une section de `string` délimité par des espace ou par le début/fin d'une `string`.

-   L'output sera suivi d'un retour à la ligne (`'\n'`).

-   Si le nombre de paramètres n'est pas 1, ou si il n'y a pas de mots, le programme affiche un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "hello there"
hello
student@ubuntu:~/piscine-go/test$ ./test "hello   .........  bye"
hello
student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$ ./test "hello" "there"

student@ubuntu:~/piscine-go/test$
```
