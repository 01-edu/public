## uniqueoccurences

### Instructions

Write a program that outputs `true` if the number of occurrences of each character is unique, otherwise `false`. `\n` should be at the end of the line.

If number of arguments is not 1 output nothing.

Only lower case characters will be given.

### Usage

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./main abbaac
true
student@ubuntu:~/[[ROOT]]/test$ ./main ab
false
student@ubuntu:~/[[ROOT]]/test$ ./main abcacccazb
true
```

In first example, character 'a' has 3 occurrences, 'b' has 2 and 'c' has 1. No two characters have the same number of occurrences.
