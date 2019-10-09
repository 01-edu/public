## brainfuck

### Instructions

Écrire un program interpréteur du `Brainfuck`.
Le code source sera donné en premier paramètre.
Le code sera toujours valide, avec moins de 4096 opérations.
Le `Brainfuck` est un langage minimaliste. Il consiste en un slice de `byte` (octets) (dans cet exercise 2048 octets) tous initialisés à zéro, et avec un pointeur sur le premier octet.

Chaque opérateur consiste en un seul caractère :

-   '>' incrémente le pointeur
-   '<' décrémente le pointeur
-   '+' incrémente le byte pointé
-   '-' décrémente le byte pointé
-   '.' affiche le byte pointé sur la sortie standard
-   '[' se rend à son ']' correspondant si le byte pointé est 0 (début de la boucle)
-   ']' se rend à son '[' correspondant si le byte pointé n'est pas 0 (fin de la boucle)

Tout autre caractère est un commentaire.

### Utilisation

```console
student@ubuntu:~/piscine-go/brainfuck$ go build
student@ubuntu:~/piscine-go/brainfuck$ ./brainfuck "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>." | cat -e
Hello World!$
student@ubuntu:~/piscine-go/brainfuck$ ./brainfuck "+++++[>++++[>++++H>+++++i<<-]>>>++\n<<<<-]>>--------.>+++++.>." | cat -e
Hi$
student@ubuntu:~/piscine-go/brainfuck$ ./brainfuck "++++++++++[>++++++++++>++++++++++>++++++++++<<<-]>---.>--.>-.>++++++++++." | cat -e
abc$
student@ubuntu:~/piscine-go/brainfuck$ ./brainfuck | cat -e
$
student@ubuntu:~/piscine-go/brainfuck$
```
