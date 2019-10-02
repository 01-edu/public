## range

### Instructions

Write the program which must:

- **Allocate (with make())** an array of integers.

- Fill it with consecutive values that begin at the first argument and end at  the second argument (Including first and second !).

- That prints the array.

- In case of error you should handle it.

- And if the number of arguments is bigget or lower than 2 it should print `\n`.

### Expected output :

```console
student@ubuntu:~/range$ go build
student@ubuntu:~/range$ ./range 1 3
[1 2 3]
student@ubuntu:~/range$ ./range -1 2
[-1 0 1 2]
student@ubuntu:~/range$ ./range 0 0
[0]
student@ubuntu:~/range$
```
