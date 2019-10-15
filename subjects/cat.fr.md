## cat

### Instructions

Écrire un programme qui a le même comportement que la ligne de commande `cat`.

-   Les `options` ne doivent pas être gérés.

-   Si le programme est éxécuté sans arguments il doit prendre l'`input` et l'afficher.

-   Dans le dossier du programme créer deux fichiers nommés `quest8.txt` et `quest8T.txt`.

-   Copier dans le fichier `quest8.txt` la phrase suivante :

`"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing`

-   Copier dans le fichier `quest8T.txt` la phrase suivante :

`"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."`

-   En cas d'erreur le programme doit imprimer l'erreur.

-   L eprogramme doit être rendu dans un dossier nommé `cat`.

### Utilisation:

```console
student@ubuntu:~/piscine-go/cat$ go build
student@ubuntu:~/piscine-go/cat$ ./cat
Hello
Hello
student@ubuntu:~/piscine-go/cat$ ./cat quest8.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
student@ubuntu:~/piscine-go/cat$ ./cat quest8.txt quest8T.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing

"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."
```
