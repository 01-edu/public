## piglatin

### Instructions

Write a **program** that transforms a string passed as argument in its `Pig Latin` version.

The rules used by Pig Latin are as follows:

- If a word begins with a vowel, just add "ay" to the end.
- If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.

If the word has no vowels the program should print "No vowels".

If the number of arguments is different from one, the program prints a newline ("`\n`").

### Usage

```console
student@ubuntu:~/student/test$ go build
student@ubuntu:~/student/test$ ./piglatin

student@ubuntu:~/student/test$ ./piglatin pig | cat -e
igpay$
student@ubuntu:~/student/test$ ./piglatin Is | cat -e
Isay$
student@ubuntu:~/student/test$ ./piglatin crunch | cat -e
unchcray$
student@ubuntu:~/student/test$ ./piglatin crnch | cat -e
No vowels$
student@ubuntu:~/student/test$ ./piglatin something else | cat -e
$
```
