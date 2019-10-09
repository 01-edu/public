## itoabaseprog

### Instructions

Écrire un programme qui:

- convertit un nombre en base 10 en nombre dans la base spécifiée
- reçoit deux paramètres:
  - le premier est la valeur
  - le deuxième est la base

Cette base est exprimée comme un `int`, de 2 à 16. Les caractères compris dans la base sont les chiffres de 0 à 9, suivis des lettres majuscules de A à F.

Par exemple, la base `4` sera équivalente à "0123" et la base `16` sera équivalente à "0123456789ABCDEF".

Si la valeur est négative, la `string` résultante doit être précédée d'un signe moins `-`.

### Utilisation

```console
student@ubuntu:~/piscine-go/itoabaseprog$ go build
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 15 16
F
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 255 2
11111111
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog -4 2 | cat -e
-100$
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog -df sdf
The value "-df" can not be converted to an int
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 23 ew | cat -e
The value "ew" can not be converted to an int$
student@ubuntu:~/piscine-go/itoabaseprog$ ./itoabaseprog 4 23

student@ubuntu:~/piscine-go/itoabaseprog$ 323 12
22B
student@ubuntu:~/piscine-go/itoabaseprog$
```
