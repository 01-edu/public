## searchreplace

### Instructions

Écrire un programme qui prends 3 arguments, le premier argument est une `string` dans laquelle une lettre (le 2éme argument) est remplacée par une autre (3éme argument).

-   Si le nombre d'aruments n'est pas 3, le programme affiche un retour à la ligne (`'\n'`).

-   Si le second argument n'est pas contenu dans le premier (la `string`) alors le programme réécris la `string` suivi d'un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "hella there" "a" "o"
hello there
student@ubuntu:~/piscine-go/test$ ./test "abcd" "z" "l"
abcd
student@ubuntu:~/piscine-go/test$ ./test "something" "a" "o" "b" "c"

student@ubuntu:~/piscine-go/test$
```
