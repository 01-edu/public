## quadChecker

### Instructions

This raid is based on the `quad` functions.

Create a program `quadChecker` that takes a `string` as an argument and displays the name of the matching `quad` and its dimensions.

- If the argument is not a `raid` the program should print `Not a Raid function`.

- All answers must end with a newline (`'\n'`).

- If there is more than one `quad` matches, the program must display them all alphabetically ordered and separated by a `||`.

### Usage

- If it's `quadA`

```console
student@ubuntu:~/[[ROOT]]/quadChecker$ ls -l
-rw-r--r-- 1 student student  nov 23 14:30 main.go
-rwxr-xr-x 1 student student  nov 23 19:18 quadChecker
-rwxr-xr-x 1 student student  nov 23 19:50 quadA
-rwxr-xr-x 1 student student  nov 23 19:50 quadB
-rwxr-xr-x 1 student student  nov 23 19:50 quadC
-rwxr-xr-x 1 student student  nov 23 19:50 quadD
-rwxr-xr-x 1 student student  nov 23 19:50 quadE
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadA 3 3 | ./quadChecker
[quadA] [3] [3]
student@ubuntu:~/[[ROOT]]/quadChecker$
student@ubuntu:~/[[ROOT]]/quadChecker$
student@ubuntu:~/[[ROOT]]/quadChecker$
student@ubuntu:~/[[ROOT]]/quadChecker$
```

- If it's `quadC 1 1` :

```console
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadC 1 1
A
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadD 1 1
A
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadE 1 1
A
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadC 1 1 | ./quadChecker
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
student@ubuntu:~/[[ROOT]]/quadChecker$
```

- If it's `quadC 1 2` :

```console
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadE 1 2
A
C
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadC 1 2
A
C
student@ubuntu:~/[[ROOT]]/quadChecker$ ./quadE 1 2 | ./quadChecker
[quadC] [1] [2] || [quadE] [1] [2]
student@ubuntu:~/[[ROOT]]/quadChecker$
```
