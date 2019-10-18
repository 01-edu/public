## raid3

### Instructions

This raid is based on the `raid1` functions.

Create a program `raid3` that takes a `string` as an argument and displays the name of the matching `raid1` and its dimensions.

-   If the argument is not a `raid1` the program should print `Not a Raid function`.

-   All answers must end with a newline (`'\n'`).

-   If there is more than one `raid1` matches, the program must display them all alphabetically ordered and separated by a `|`.

### Usage

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ echo HELLO | ./raid3
Not a Raid function
student@ubuntu:~/piscine-go/test$ ./raid1a 4 4
o--o
|  |
|  |
o--o
student@ubuntu:~/piscine-go/test$ ./raid1a 4 4 | ./raid3
[raid1a] [4] [4]
student@ubuntu:~/piscine-go/test$ ./raid1a 3 4 | ./raid3
[raid1a] [3] [4]
student@ubuntu:~/piscine-go/test$ ./raid1c 1 1
A
student@ubuntu:~/piscine-go/test$ ./raid1d 1 1
A
student@ubuntu:~/piscine-go/test$ ./raid1e 1 1
A
student@ubuntu:~/piscine-go/test$ ./raid1c 1 1 | ./raid3
[raid1c] [1] [1] || [raid1d] [1] [1] || [raid1e] [1] [1]
student@ubuntu:~/piscine-go/test$
```
