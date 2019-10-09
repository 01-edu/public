## revWstr

### Instructions

Écrire un programme qui prend une `string` comme paramètre, et affiche ses mots en sens inverse.

-   Un mot est une suite de caractères **alphanumériques.**

-   Si le nombre de paramètres est différent de 1, le programme affiche un retour à la ligne (`'\n'`).

-   Dans les paramètres qui seront testés, il n'y aura pas d'espaces extra. (ce qui signifie qu'il n'y aura pas d'espaces additionnels, ni au début, ni à la fin de la `string` et que les mots seront toujours séparés par un seul espace).

### Utilisation

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test "the time of contempt precedes that of indifference"
indifference of that precedes contempt of time the
student@ubuntu:~/piscine-go/test$ ./test "abcdefghijklm"
abcdefghijklm
student@ubuntu:~/piscine-go/test$ ./test "he stared at the mountain"
mountain the at stared he
student@ubuntu:~/piscine-go/test$ ./test

student@ubuntu:~/piscine-go/test$
```
