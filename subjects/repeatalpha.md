# repeatalpha
## Instructions

Write a program called repeat_alpha that takes a string and display it
repeating each alphabetical character as many times as its alphabetical index,
followed by a newline.

'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

Case remains unchanged.

If the number of arguments is not 1, just display a newline.

Examples:

```console
student@ubuntu:~/student/repeatalpha$ go build
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "abc" | cat -e
abbccc
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "Alex." | cat -e
Alllllllllllleeeeexxxxxxxxxxxxxxxxxxxxxxxx.$
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "abacadaba 42!" | cat -e
abbacccaddddabba 42!$
student@ubuntu:~/student/repeatalpha$ ./repeatalpha | cat -e
$
student@ubuntu:~/student/repeatalpha$ ./repeatalpha "" | cat -e
$
student@ubuntu:~/student/repeatalpha$ 
```
