## capitalizeprog

### Instructions

Écrire un programme qui met en majuscule la première lettre de chaque mot et en minuscule les autres lettres du reste du mot d'une `string`.

-   Un mot est une suite de caractères **alphanumériques**.

-   Si il y a plus d'un argument le programme doit afficher `Too many arguments`.

-   Si il n'y a pas d'arguments le programme doit afficher un retour à la ligne (`'\n'`).

### Utilisation :

```console
student@ubuntu:~/capitalizeprog$ go build
student@ubuntu:~/capitalizeprog$ ./capitalizeprog "Hello! How are you? How+are+things+4you?" | cat -e
Hello! How Are You? How+Are+Things+4you?$
student@ubuntu:~/capitalizeprog$ ./capitalizeprog Hello! How are you? | cat -e
Too many arguments$
student@ubuntu:~/capitalizeprog$ ./capitalizeprog

student@ubuntu:~/capitalizeprog$
```
