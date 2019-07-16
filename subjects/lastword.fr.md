## lastword

### Instructions

Écrire un programme qui prend une `string` et qui affiche son dernier mot, suivi d'un newline(`'\n'`).

- Un mot est une section de `string` délimitée par des espaces ou par le début/fin d'une `string`.

- L'output sera suivi d'un newline(`'\n'`).

- Si le nombre de paramètres est différent de 1, ou si il n'y a pas de mot, le programme affiche un newline(`'\n'`).

### Utilisation

```console
student@ubuntu:~/student/lastword$ go build
student@ubuntu:~/student/lastword$ ./lastword "FOR PONY" | cat -e
PONY$
student@ubuntu:~/student/lastword$ ./lastword "this        ...       is sparta, then again, maybe    not" | cat -e
not$
student@ubuntu:~/student/lastword$ ./lastword "  " | cat -e
$
student@ubuntu:~/student/lastword$ ./lastword "a" "b" | cat -e
$
student@ubuntu:~/student/lastword$ ./lastword "  lorem,ipsum  " | cat -e
lorem,ipsum$
student@ubuntu:~/student/lastword$
```
