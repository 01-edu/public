## hiddenp

### Instructions

Write a program named `hiddenp` that takes two `string` and that, if the first `string` is hidden in the second one, displays `1` followed by a newline (`'\n'`), otherwise it displays `0` followed by a newline.

Let s1 and s2 be `string`. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, **in the same order as they appear in s1.**

If s1 is an empty `string` it is considered hidden in any `string`.

If the number of parameters is different from 2, the program displays a newline.

### Usage

```console
student@ubuntu:~/piscine-go/hiddenp$ go build
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
1$
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp "abc" "2altrb53c.sse" | cat -e
1$
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp "abc" "btarc" | cat -e
0$
student@ubuntu:~/piscine-go/hiddenp$ ./hiddenp | cat -e
$
student@ubuntu:~/piscine-go/hiddenp$
```
