## printcombn

### Instructions

-   Write a function that prints all possible combinations of **n** different digits in ascending order.

-   n will be defined as : 0 < n < 10

below are your references for the **printing format** expected.

-   (for n = 1) '0, 1, 2, 3, ...8, 9'

-   (for n = 3) '012, 013, 014, 015, 016, 017, 018, 019, 023,...689, 789'

### Expected function

```go
func PrintCombN(n int) {

}
```

### Usage

Here is a possible [program](TODO-LINK) to test your function :

```go
package main

import piscine ".."

func main() {
	piscine.PrintCombN(1)
	piscine.PrintCombN(2)
	piscine.PrintCombN(9)
}
```

And its output :

```console
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
0, 1, 2, 3, 4, 5, 6, 7, 8, 9
012, 013, 014, 015, 016, 017, 018, ... 679, 689, 789
012345678, 012345679, ..., 123456789
student@ubuntu:~/piscine-go/test$
```
