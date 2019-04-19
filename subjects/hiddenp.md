# hiddenp
## Instructions

Write a program named hidenp that takes two strings and displays 1
followed by a newline if the first string is hidden in the second one,
otherwise displays 0 followed by a newline.

Let s1 and s2 be strings. We say that s1 is hidden in s2 if it's possible to
find each character from s1 in s2, in the same order as they appear in s1.
Also, the empty string is hidden in any string.

If the number of parameters is not 2, the program displays a newline.

And its output :

```console
student@ubuntu:~/student/hiddenp$ go build
student@ubuntu:~/student/hiddenp$ ./hiddenp "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
student@ubuntu:~/student/hiddenp$ ./hiddenp "abc" "2altrb53c.sse" | cat -e
1$
student@ubuntu:~/student/hiddenp$ ./hiddenp "abc" "btarc" | cat -e
0$
student@ubuntu:~/student/hiddenp$ ./hiddenp | cat -e
$
student@ubuntu:~/student/hiddenp$
```
