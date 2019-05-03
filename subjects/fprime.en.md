## fprime

### Instructions

Write a program that takes a positive `int` and displays its prime factors on the standard output, followed by a newline.

- Factors must be displayed in ascending order and separated by `*`, so that the expression in the output gives the right result.

- If the number of parameters is not 1, simply display a newline.

- The input, when there's one, will be valid.

Example of output :

```console
student@ubuntu:~/piscine/test$ go build
student@ubuntu:~/piscine/test$ ./test 225225
3*3*5*5*7*11*13
student@ubuntu:~/piscine/test$ ./test 8333325
3*3*5*5*7*11*13*37
student@ubuntu:~/piscine/test$ ./test 9539
9539
student@ubuntu:~/piscine/test$ ./test 804577
804577
student@ubuntu:~/piscine/test$ ./test 42
2*3*7
student@ubuntu:~/piscine/test$ ./test

student@ubuntu:~/piscine/test$
```