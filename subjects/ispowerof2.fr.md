## ispowerof2

### Instructions

Écrire un programme qui détermine si un nombre donné est une puissance de 2.

Ce programme doit afficher `true` si le nombre donné est une puissance de 2, autrement il affiche `false`.

-   Si il y a plus d'un ou aucun argument le programme doit afficher un retour à la ligne (`'\n'`).

-   Les cas d'erreurs doivent être gérés.

### Utilisation :

```console
student@ubuntu:~/ispowerof2$ go build
student@ubuntu:~/ispowerof2$ ./ispowerof2 2 | cat -e
true$
student@ubuntu:~/ispowerof2$ ./ispowerof2 64 | cat -e
true$
student@ubuntu:~/ispowerof2$ ./ispowerof2 513 | cat -e
false$
student@ubuntu:~/ispowerof2$ ./ispowerof2

student@ubuntu:~/ispowerof2$ ./ispowerof2 64 1024

student@ubuntu:~/ispowerof2$ ./ispowerof2 notanumber | cat -e
strconv.Atoi: parsing "a": invalid syntax$
student@ubuntu:~/ispowerof2$
```
