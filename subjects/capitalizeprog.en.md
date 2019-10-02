## capitalizeprog

### Instructions

Write a program that capitalizes the first letter of each word **and** lowercases the rest of each word of a `string`.

- A word is a sequence of **alphanumerical** characters.

- If the number of arguments is bigger than one it should print `To many arguments`.

- If there's no arguments given to the program it should print a new line `\n`.

### Expected output :

```console
student@ubuntu:~/capitalizeprog$ go build
student@ubuntu:~/capitalizeprog$ ./capitalizeprog "Hello! How are you? How+are+things+4you?"
Hello! How Are You? How+Are+Things+4you?
student@ubuntu:~/capitalizeprog$ ./capitalizeprog Hello! How are you?
To many arguments
student@ubuntu:~/capitalizeprog$ ./capitalizeprog

student@ubuntu:~/capitalizeprog$
```
