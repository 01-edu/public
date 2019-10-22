## flags

### Instructions

Write a **program** that can have as arguments `--insert` (or `-i`),  `--order` (or `-o`) and a `string`.

This program should :

- Insert the string given to the `--insert` (or `-i`), in the `string` argument, if given.
- Order the `string` argument (in ASCII order) if given the flag `--order` (or `-o`).
- In case there are no arguments or the flag `--help` (or `-h`) is given, it should print the options, as shown in the example.

Example of output :

```console
student@ubuntu:~/piscine-go/flags$ go build
student@ubuntu:~/piscine-go/flags$ ./flags --insert=4321 --order asdad
1234aadds
student@ubuntu:~/piscine-go/flags$ ./flags --insert=4321 asdad
asdad4321
student@ubuntu:~/piscine-go/flags$ ./flags asdad
asdad
student@ubuntu:~/piscine-go/flags$ ./flags --order 43a21
1234a
student@ubuntu:~/piscine-go/flags$ ./flags
--insert
  -i
    This flag inserts the string into the string passed as argument.
--order
  -o
    This flag will behave like a boolean, if it is called it will order the argument.
student@ubuntu:~/piscine-go/flags$
student@ubuntu:~/piscine-go/flags$ ./flags -h
--insert
  -i
    This flag inserts the string into the string passed as argument.
--order
  -o
    This flag will behave like a boolean, if it is called it will order the argument.
student@ubuntu:~/piscine-go/flags$
student@ubuntu:~/piscine-go/flags$ ./flags --help
--insert
  -i
    This flag inserts the string into the string passed as argument.
--order
  -o
    This flag will behave like a boolean, if it is called it will order the argument.
student@ubuntu:~/piscine-go/flags$
```
