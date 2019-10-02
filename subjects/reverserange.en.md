## reverserange

### Instructions

Write the program which must:

- **Allocate (with make())** an array of integers.

- Fill it with consecutive values that begins at the second argument and end at the first argument (Including the first and second argument !).

- That prints the array.

- In case of error you should handle it.

- And if the number of arguments is bigger or lower than 2 it should print `\n`.

### Expected output :

```console
student@ubuntu:~/reverserange$ go build
student@ubuntu:~/reverserange$ ./reverserange 1 3
[3 2 1]
student@ubuntu:~/reverserange$ ./reverserange -1 2
[2 1 0 -1]
student@ubuntu:~/reverserange$ ./reverserange 0 0
[0]
student@ubuntu:~/reverserange$
```