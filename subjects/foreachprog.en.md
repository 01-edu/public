## foreachprog

### Instructions

Write a program that applies a given function, that's the first argument, to a set of runes, this being the second argument.

- If the number of arguments is bigger or lower that two it should print a newline `\n`.

- These are the functions you should add to your program.

```go
func add0(nbr rune) {
	fmt.Print(string(nbr))
}

func add1(nbr rune) {
	fmt.Print(string(nbr + 1))
}

func add2(nbr rune) {
	fmt.Print(string(nbr + 2))
}

```

### Expected output :

```console
student@ubuntu:~/foreachprog$ go build
student@ubuntu:~/foreachprog$ ./foreachprog add1 123456
234567
student@ubuntu:~/foreachprog$ ./foreachprog add2 123456
345678
student@ubuntu:~/foreachprog$ ./foreachprog add2

student@ubuntu:~/foreachprog$ ./foreachprog add1 1234 123

student@ubuntu:~/foreachprog$
```
