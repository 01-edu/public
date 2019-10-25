## nruneprog

##**WARNING! VERY IMPORTANT!**

For this exercise a function will be tested **with the exam own main**. However the student **still needs** to submit a structured program:

This means that:

- The package needs to be named `package main`.
- The submitted code needs one declared function main(```func main()```) even if empty.
- The function main declared needs to **also pass** the `Restrictions Checker`(illegal functions tester). It is advised for the student to just empty the function main after its own testings are done.
- Every other rules are obviously the same than for a `program`.

### Instructions

Write a function that returns the nth `rune` of a `string`.

- In case of impossibilities, the function returns `0`.

### Expected function

```go
func NRune(s string, n int) rune {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import (
	"github.com/01-edu/z01"
	piscine ".."
)

func main() {
	z01.PrintRune(piscine.NRune("Hello!", 3))
	z01.PrintRune(piscine.NRune("Salut!", 2))
	z01.PrintRune(piscine.NRune("Bye!", -1))
	z01.PrintRune(piscine.NRune("Bye!", 5))
	z01.PrintRune(piscine.NRune("Ola!", 4))
	z01.PrintRune('\n')
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
la!
student@ubuntu:~/piscine-go/test$
```
