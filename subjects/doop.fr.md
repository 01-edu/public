## doop

### Instructions

Écrire un [programme](TODO-LINK) qui s'apelle `doop`.

Le programme doit être utilisé avec trois arguments:

- Une valeur
- Un opérateur
- Une autre valeur

En cas d'argument invalide le programme affiche `0`.

En cas de nombre invalide d'arguments le programme affiche rien.

`fmt.Print` est autorisé.

### Utilisation

```console
student@ubuntu:~/piscine/test$ go build doop.go
student@ubuntu:~/piscine/test$ ./doop
student@ubuntu:~/piscine/test$ ./doop 1 + 1
2
student@ubuntu:~/piscine/test$ ./doop hello + 1
0
student@ubuntu:~/piscine/test$ ./doop 1 p 1
0
student@ubuntu:~/piscine/test$ ./doop 1 + 1
2
student@ubuntu:~/piscine/test$ ./doop 1 / 0
No division by 0
student@ubuntu:~/piscine/test$ ./doop 1 % 0
No modulo by 0
student@ubuntu:~/piscine/test$ ./doop 1 * 1
1

```
