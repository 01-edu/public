## Cat

### Instructions

Write a program that does the same thing as the system's `cat` command-line.

- You don't have to handle options.

- But if just call the program with out arguments it should take a input and print it back

- In the program folder create two files named `quest8.txt` and `quest8T.txt`.

- Copy to the `quest8.txt` file this :

  - "Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing

- Copy to the `quest8T.txt` file this :

  - "Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."

- In case of error it should print the error.

### Output:

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./test
Hello
Hello
student@ubuntu:~/student/test$ ./test quest8.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
student@ubuntu:~/student/test$ ./test quest8.txt quest8T.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing

"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."

```
