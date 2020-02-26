## reverserange

### Instructions

Écrire un programme qui doit:

- Allouer (avec `make`) une slice d'entiers.

- Le remplir avec des valeurs consécutives qui commencent au deuxième argument et qui finissent au premier (en incluant les valeurs des deux arguments !)

- Et qui affiche cette slice.

Les erreurs doivent être gérées.

Si le nombre d'arguments est différent de 2 le programme affiche un retour à la ligne (`'\n'`).

### Utilisation :

```console
student@ubuntu:~/reverserange$ go build
student@ubuntu:~/reverserange$ ./reverserange 1 3
[3 2 1]
student@ubuntu:~/reverserange$ ./reverserange -1 2 | cat -e
[2 1 0 -1]$
student@ubuntu:~/reverserange$ ./reverserange 0 0
[0]
student@ubuntu:~/reverserange$ ./reverserange 0 -3
[-3 -2 -1 0]
student@ubuntu:~/reverserange$ ./reverserange 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
student@ubuntu:~/reverserange$
```
