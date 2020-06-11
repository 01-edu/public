## printstr

### Instructions

- Write a function that prints one by one the characters of a `string` on the screen.

### Expected function

```go
func PrintStr(s string) {

}
```

### Hints

Here is a possible program to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.PrintStr("Hello World!")
}
```

And its output :

```console
student@ubuntu:~/[[ROOT]]/test$ go build
student@ubuntu:~/[[ROOT]]/test$ ./test | cat -e
Hello World!%
student@ubuntu:~/[[ROOT]]/test$
```
