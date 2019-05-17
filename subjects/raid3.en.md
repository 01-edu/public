## raid3

### Instructions

Creare a program that takes a character string as argument and displays the name of the raid in question, as well as its dimensions.

- Executable name : `raid3`.

- If the argument isn't a raid should print `Not a Raid function`.

Example of this output :

```console
student@ubuntu:~/studen/test$ go build
student@ubuntu:~/studen/test$ echo HELLO | ./raid3
Not a Raid function
student@ubuntu:~/studen/test$
```

- Whatever the answer, your line must be ended by a `\n`.

- If there is more than one raid1, you must display them all alphabetivally.

Examples of output :

```console
student@ubuntu:~/studen/test$ ./raid1a 4 4
o--o
|  |
|  |
o--o
student@ubuntu:~/studen/test$ ./raid1a 4 4 | ./raid3
[raid1a] [4] [4]
student@ubuntu:~/studen/test$ ./raid1a 3 4 | ./raid3
[raid1a] [3] [4]
student@ubuntu:~/studen/test$ ./raid1c 1 1
A
student@ubuntu:~/studen/test$ ./raid1d 1 1
A
student@ubuntu:~/studen/test$ ./raid1e 1 1
A
student@ubuntu:~/studen/test$ ./raid1c 1 1 | ./raid3
[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]
student@ubuntu:~/studen/test$
```