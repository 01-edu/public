## range

### Instructions

Écrire un programme qui doit:

- Allouer (avec `make`) une slice d'entiers.

- Le remplir avec des valeurs consécutives qui commencent au premier argument et qui finissent au deuxième (En incluant les valeurs des deux arguments !)

- Et qui affiche cette slice.

Les erreurs doivent être gérées.

Si le nombre d'arguments est différent de 2 le programme affiche un retour à la ligne (`'\n'`).

### Utilisation

```console
student@ubuntu:~/range$ go build
student@ubuntu:~/range$ ./range 1 3
[1 2 3]
student@ubuntu:~/range$ ./range -1 2 | cat -e
[-1 0 1 2]$
student@ubuntu:~/range$ ./range 0 0
[0]
student@ubuntu:~/reverserange$ ./reverserange 0 nan | cat -e
strconv.Atoi: parsing "nan": invalid syntax$
student@ubuntu:~/range$
```
