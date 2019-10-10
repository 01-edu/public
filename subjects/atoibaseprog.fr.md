## atoibaseprog

### Instructions

Écrire un programme qui prend un nombre `string` et sa base `string` en paramètres et retourne sa conversion en `int`.

- Si la base ou le nombre `string` n'est pas valide le programme retourne `0`:

- Si le nombre d'argument est différent de deux alors le programme affiche un newline ("`\n`").

Règles de validité d'une base :

- Une base doit contenir au moins 2 caractères.
- Chaque caractère d'une base doit être unique.
- Une base ne doit pas contenir les caractères `+` ou `-`.

Seuls des nombres en `string` valides seront testés.

La fonction **ne doit pas** gérer les nombres négatifs.

### Expected output :

```console
student@ubuntu:~/atoibaseprog$ go build
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 125 0123456789
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 1111101 01
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 7D 0123456789ABCDEF
125
student@ubuntu:~/atoibaseprog$ ./atoibaseprog uoi choumi | cat -e
125$
student@ubuntu:~/atoibaseprog$ ./atoibaseprog bbbbbab -ab | cat -e
0$
student@ubuntu:~/atoibaseprog$ ./atoibaseprog 1111101

student@ubuntu:~/atoibaseprog$
```
