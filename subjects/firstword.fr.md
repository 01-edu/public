## firstword

### Instructions

Écrire un programme qui prend une `string` et qui affiche son premier mot, suivi d'un newline(`'\n'`).

- Un mot est une section de `string` délimité par des espace ou par le début/fin d'une `string`.

- L'output sera suivi d'un newline(`'\n'`).

- Si le nombre de paramètres n'est pas 1, ou si il n'y a pas de mots, le programme affiche un newline(`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test "hello there"
hello
student@ubuntu:~/piscine/test$ ./test "hello   .........  bye"
hello
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$ ./test "hello" "there"

student@ubuntu:~/piscine/test$
```
