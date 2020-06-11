## raid3

### Instructions

This raid is based on the `raid1` functions.

Create a program `raid3` that takes a `string` as an argument and displays the name of the matching `raid1` and its dimensions.

- If the argument is not a `raid1` the program should print `Not a Raid function`.

- All answers must end with a newline (`'\n'`).

- If there is more than one `raid1` matches, the program must display them all alphabetically ordered and separated by a `||`.

### Usage

- If it's `raid1a`

```console
student@ubuntu:~/[[ROOT]]/raid3$ ls -l
-rw-r--r-- 1 student student  nov 23 14:30 main.go
-rwxr-xr-x 1 student student  nov 23 19:18 raid3
-rwxr-xr-x 1 student student  nov 23 19:50 raid1a
-rwxr-xr-x 1 student student  nov 23 19:50 raid1b
-rwxr-xr-x 1 student student  nov 23 19:50 raid1c
-rwxr-xr-x 1 student student  nov 23 19:50 raid1d
-rwxr-xr-x 1 student student  nov 23 19:50 raid1e
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1a 3 3 | ./raid3
[raid1a] [3] [3]
student@ubuntu:~/[[ROOT]]/raid3$
student@ubuntu:~/[[ROOT]]/raid3$
student@ubuntu:~/[[ROOT]]/raid3$
student@ubuntu:~/[[ROOT]]/raid3$
```

- If it's `raidc 1 1` :

```console
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1c 1 1
A
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1d 1 1
A
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1e 1 1
A
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1c 1 1 | ./raid3
[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]
student@ubuntu:~/[[ROOT]]/raid3$
```

- If it's `raidc 1 2` :

```console
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1e 1 2
A
C
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1c 1 2
A
C
student@ubuntu:~/[[ROOT]]/raid3$ ./raid1e 1 2 | ./raid3
[raid1c] [1] [2] || [raid1e] [1] [2]
student@ubuntu:~/[[ROOT]]/raid3$
```
