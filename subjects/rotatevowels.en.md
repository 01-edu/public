## rotatevowels

### Instructions

Write a **program** that checks the arguments for vowels.

- If the argument contains vowels you have to reverse the order of the vowels in the argument.
- If the number of arguments are less than 1, the program display a new line, `\n`.
- If the arguments doesn't have any vowels, the program just print the arguments.

Example of output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./rotatevowels "Hello World"
Hollo Werld
student@ubuntu:~/piscine-go/test$ ./rotatevowels "Hello World" "problem solved"
Hello Werld problom solved
student@ubuntu:~/piscine-go/test$ ./rotatevowels "str" "shh" "psst"
str shh psst
student@ubuntu:~/piscine-go/test$ ./rotatevowels "happy thoughts" "good luck"
huppy thooghts guod lack
student@ubuntu:~/piscine-go/test$ ./rotatevowels

student@ubuntu:~/piscine-go/test$
```
