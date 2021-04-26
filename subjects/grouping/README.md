## grouping

### Instructions

Write a program that receives two strings and replicates the use of brackets in regular expressions. Brackets in regular expressions returns the words that contain the expression inside of it.

The program should handle the "`|`" operator, that searches for both strings on each side of the operator.

The output of the program should be the results of the regular expression by order of appearance in the string, being themselves identified by a number.

If the number of arguments is different from 2, if the regular expression is not valid, if the last argument is empty or there are no matches the program should print nothing.

### Usage

```console
student@ubuntu:~/grouping/test$ go build
student@ubuntu:~/grouping/test$ ./regbrackets "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
1: heavy
2: steady
3: heavy
student@ubuntu:~/grouping/test$ ./regbrackets "(e|n)" "I currently have 4 windows opened up… and I don’t know why."
1: currently
2: currently
3: have
4: windows
5: opened
6: opened
7: and
8: don’t
9: know
student@ubuntu:~/grouping/test$ ./regbrackets "(hi)" "He swore he just saw his sushi move."
1: his
2: sushi
student@ubuntu:~/grouping/test$ ./regbrackets "(s)" ""
student@ubuntu:~/grouping/test$ ./regbrackets "i" "Something in the air"
student@ubuntu:~/grouping/test$
```
