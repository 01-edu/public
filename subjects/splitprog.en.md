## splitprog

### Instructions

Write a program that separates the words of a `string` and puts them in a `string` array and then prints it to standard output.

- The program receives two parameters:

  - The first parameter is the `string`
  - The second parameter is the `separator`

### Expected output :

```console
student@ubuntu:~/piscine-go/splitprog$ go build
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "HelloHAhowHAareHAyou?" HA
[Hello how are you?]
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "Hello,how,are,you?" ","
[Hello how are you?]
student@ubuntu:~/piscine-go/splitprog$ ./splitprog "HelloHAhowHAareHAyou?"
student@ubuntu:~/piscine-go/splitprog$
student@ubuntu:~/piscine-go/splitprog$ ./splitprog
student@ubuntu:~/piscine-go/splitprog$
```
