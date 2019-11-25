## cat

### Instructions

Write a program that has the same behaviour as the system's `cat` command-line.

-   The `options` do not have to be handled.

-   If the program is called without arguments it should take the `input` and print it back.

-   In the program folder create two files named `quest8.txt` and `quest8T.txt`.

-   Copy to the `quest8.txt` file the following sentence :

`"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing`

-   Copy to the `quest8T.txt` file the following sentence :

`"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."`

-   In case of error the program should print the error.

-   The program must be submitted inside a folder named `cat`.

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
