## atoiprog

### Instructions

- Write a program that simulates the behaviour of the `Atoi` function in Go. `Atoi` transforms a number represented as a `string` in a number represented as an `int` and after that subtract 1.

- `Atoi` returns `0` if the `string` is not considered as a valid number. For this exercise **non-valid `string` chains will be tested**. Some will contain non-digits characters.

- For this exercise the handling of the signs + or - **does have** to be taken into account.

- See the example for more details

### Expected output :

```console
student@ubuntu:~/piscine-go/atoiprog$ go build
student@ubuntu:~/piscine-go/atoiprog$ ./atoiprog 12345
12344
student@ubuntu:~/piscine-go/atoiprog$ ./atoiprog 32f
-1
student@ubuntu:~/piscine-go/atoiprog$ ./atoiprog jfasl
-1
student@ubuntu:~/piscine-go/atoiprog$ ./atoiprog -3
-4
student@ubuntu:~/piscine-go/atoiprog$
```
