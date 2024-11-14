## cat

### Instructions

Write a program that behaves like a simplified `cat` command.

- The options do not have to be handled.

- If the program is called without arguments it should take the standard input (stdin) and print it back on the standard output (stdout).

```console
$ echo '"Programming is a skill best acquired by practice and example rather than from books" by Alan Turing' > quest8.txt
$ cat <<EOF> quest8T.txt
"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."
EOF
$ go run . abc
ERROR: open abc: no such file or directory
exit status 1
$ go run . quest8.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
$ go run . quest8.txt abc
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
ERROR: open abc: No such file or directory
$ cat quest8.txt | go run .
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
$ go run .
Hello
Hello
^C
$ go run . quest8.txt quest8T.txt
"Programming is a skill best acquired by pratice and example rather than from books" by Alan Turing
"Alan Mathison Turing was an English mathematician, computer scientist, logician, cryptanalyst. Turing was highly influential in the development of theoretical computer science, providing a formalisation of the concepts of algorithm and computation with the Turing machine, which can be considered a model of a general-purpose computer. Turing is widely considered to be the father of theoretical computer science and artificial intelligence."
$
```
