# alphamirror
## Instructions

Write a program called alpha_mirror that takes a string and displays this string
after replacing each alphabetical character by the opposite alphabetical
character, followed by a newline.

'a' becomes 'z', 'Z' becomes 'A'
'd' becomes 'w', 'M' becomes 'N'

and so on.

Case is not changed.

If the number of arguments is not 1, display only a newline.

And its output :

```console
student@ubuntu:~/student/alphamirror$ go build
student@ubuntu:~/student/alphamirror$ ./alphamirror "abc"
zyx
student@ubuntu:~/student/alphamirror$ ./alphamirror "My horse is Amazing." | cat -e
Nb slihv rh Znzarmt.$
student@ubuntu:~/student/alphamirror$ ./alphamirror | cat -e
$
student@ubuntu:~/student/alphamirror$ 
```
